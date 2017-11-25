// Some functions in this file are from "github.com/go-playground/pure".

package http2

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"strings"
)

var (
	xmlHeaderBytes = []byte(xml.Header)
)

func detectContentType(filename string) (t string) {
	if t = mime.TypeByExtension(filepath.Ext(filename)); t == "" {
		t = OctetStream
	}
	return
}

// AcceptedLanguages returns an array of accepted languages denoted by
// the Accept-Language header sent by the browser
// NOTE: some stupid browsers send in locales lowercase when all the rest send it properly
func AcceptedLanguages(r *http.Request) (languages []string) {
	var accepted string
	if accepted = r.Header.Get(AcceptedLanguage); accepted == blank {
		return
	}

	options := strings.Split(accepted, ",")
	l := len(options)
	languages = make([]string, l)
	for i := 0; i < l; i++ {
		locale := strings.SplitN(options[i], ";", 2)
		languages[i] = strings.Trim(locale[0], " ")
	}

	return
}

// Attachment is a helper method for returning an attachement file
// to be downloaded, if you with to open inline see function Inline
func Attachment(w http.ResponseWriter, r io.Reader, filename string) (err error) {
	w.Header().Set(ContentDisposition, "attachment;filename="+filename)
	w.Header().Set(ContentType, detectContentType(filename))
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, r)
	return
}

// Inline is a helper method for returning a file inline to
// be rendered/opened by the browser
func Inline(w http.ResponseWriter, r io.Reader, filename string) (err error) {
	w.Header().Set(ContentDisposition, "inline;filename="+filename)
	w.Header().Set(ContentType, detectContentType(filename))
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, r)
	return
}

// ClientIP implements a best effort algorithm to return the real client IP, it parses
// X-Real-IP and X-Forwarded-For in order to work properly with reverse-proxies such us: nginx or haproxy.
func ClientIP(r *http.Request) (clientIP string) {
	var values []string

	if values, _ = r.Header[XRealIP]; len(values) > 0 {

		clientIP = strings.TrimSpace(values[0])
		if clientIP != blank {
			return
		}
	}

	if values, _ = r.Header[XForwardedFor]; len(values) > 0 {
		clientIP = values[0]

		if index := strings.IndexByte(clientIP, ','); index >= 0 {
			clientIP = clientIP[0:index]
		}

		clientIP = strings.TrimSpace(clientIP)
		if clientIP != blank {
			return
		}
	}

	clientIP, _, _ = net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))

	return
}

// JSON marshals provided interface + returns JSON + status code
func JSON(w http.ResponseWriter, status int, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return JSONBytes(w, status, b)
}

// JSONBytes returns provided JSON response with status code
func JSONBytes(w http.ResponseWriter, status int, b []byte) (err error) {
	w.Header().Set(ContentType, ApplicationJSONCharsetUTF8)
	w.WriteHeader(status)
	_, err = w.Write(b)
	return
}

// JSONP sends a JSONP response with status code and uses `callback` to construct
// the JSONP payload.
func JSONP(w http.ResponseWriter, status int, i interface{}, callback string) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	w.Header().Set(ContentType, ApplicationJavaScriptCharsetUTF8)
	w.WriteHeader(status)

	if _, err = w.Write([]byte(callback + "(")); err == nil {
		if _, err = w.Write(b); err == nil {
			_, err = w.Write([]byte(");"))
		}
	}

	return err
}

// XML marshals provided interface + returns XML + status code
func XML(w http.ResponseWriter, status int, i interface{}) error {
	b, err := xml.Marshal(i)
	if err != nil {
		return err
	}

	return XMLBytes(w, status, b)
}

// XMLBytes returns provided XML response with status code
func XMLBytes(w http.ResponseWriter, status int, b []byte) (err error) {
	w.Header().Set(ContentType, ApplicationXMLCharsetUTF8)
	w.WriteHeader(status)

	if _, err = w.Write(xmlHeaderBytes); err == nil {
		_, err = w.Write(b)
	}

	return
}

// DecodeJSON decodes the request body into the provided struct and limits the request size via
// an io.LimitReader using the maxMemory param.
//
// The Content-Type e.g. "application/json" and http method are not checked.
func DecodeJSON(r *http.Request, maxMemory int64, v interface{}) (err error) {
	return json.NewDecoder(io.LimitReader(r.Body, maxMemory)).Decode(v)
}

// DecodeXML decodes the request body into the provided struct and limits the request size via
// an io.LimitReader using the maxMemory param.
//
// The Content-Type e.g. "application/xml" and http method are not checked.
func DecodeXML(r *http.Request, maxMemory int64, v interface{}) (err error) {
	return xml.NewDecoder(io.LimitReader(r.Body, maxMemory)).Decode(v)
}

// Decode takes the request and attempts to discover it's content type via
// the http headers and then decode the request body into the provided struct.
// Example if header was "application/json" would decode using
// json.NewDecoder(io.LimitReader(r.Body, maxMemory)).Decode(v).
//
// Notice: At present it only supports to decode JSON and XML.
func Decode(r *http.Request, maxMemory int64, v interface{}) (err error) {
	typ := r.Header.Get(ContentType)
	if idx := strings.Index(typ, ";"); idx != -1 {
		typ = typ[:idx]
	}

	switch typ {
	case ApplicationJSON:
		err = DecodeJSON(r, maxMemory, v)
	case ApplicationXML:
		err = DecodeXML(r, maxMemory, v)
	}

	return
}
