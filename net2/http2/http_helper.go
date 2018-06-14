package http2

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/xgfone/go-tools/io2"
	"github.com/xgfone/go-tools/pools"
	"github.com/xgfone/go-tools/types"
)

var (
	xmlHeaderBytes = []byte(xml.Header)

	decodeMap = make(map[string]func(*http.Request, int64, interface{}) error, 4)
)

func init() {
	RegisterDecode(ApplicationJSON, DecodeJSON)
	RegisterDecode(ApplicationXML, DecodeXML)
}

// RegisterDecode registers the decode function for Content-Type ctype.
//
// ctype is the Content-Type to decode.
//
// f is the decoder for the Content-Type ctype. The decoder has three arguments.
// The first is the request, the second is the max size of the request body,
// and the last is the decoded values, which is a pointer or map in general.
//
// If giving force as true, it will override the registered content-type decoder.
func RegisterDecode(ctype string, f func(*http.Request, int64, interface{}) error,
	force ...bool) error {
	if _, ok := decodeMap[ctype]; ok {
		if len(force) == 0 || !force[0] {
			return fmt.Errorf("the Content-Type '%s' has existed", ctype)
		}
	}
	decodeMap[ctype] = f
	return nil
}

// Decode takes the request and attempts to discover it's content type via
// the http headers and then decode the request body into the provided struct.
// Example if header was "application/json" would decode using
// json.NewDecoder(io.LimitReader(r.Body, maxMemory)).Decode(v).
//
// Notice: At present it only supports to decode JSON and XML.
func Decode(r *http.Request, maxMemory int64, v interface{}) (err error) {
	ct := GetContentType(r)
	if f := decodeMap[ct]; f != nil {
		return f(r, maxMemory, v)
	}
	return fmt.Errorf("no decoder of Content-Type %s", ct)
}

func decode(r *http.Request, maxMemory int64, v interface{},
	f func(io.Reader, interface{}) error) (err error) {

	if maxMemory <= 0 {
		maxMemory = DefaultMaxBodySize
	}
	reader := io.LimitReader(r.Body, maxMemory)

	w := pools.DefaultBufferPool.Get()
	if err = io2.ReadNWriter(w, reader, r.ContentLength); err != nil {
		pools.DefaultBufferPool.Put(w)
		return err
	}
	err = f(w, v)
	pools.DefaultBufferPool.Put(w)
	return err
}

// GeneralDecoder returns a general decoder.
func GeneralDecoder(f func(io.Reader, interface{}) error) func(*http.Request, int64, interface{}) error {
	return func(r *http.Request, maxMemory int64, v interface{}) (err error) {
		return decode(r, maxMemory, v, f)
	}
}

// DecodeJSON decodes the request body into the provided struct and limits
// the request size via an io.LimitReader using the maxMemory param.
//
// The Content-Type e.g. "application/json" and http method are not checked.
//
// If the maxMemory is equal to or less than 0, it's DefaultMaxBodySize.
func DecodeJSON(r *http.Request, maxMemory int64, v interface{}) (err error) {
	return decode(r, maxMemory, v, func(reader io.Reader, v interface{}) error {
		return json.NewDecoder(reader).Decode(v)
	})
}

// DecodeXML decodes the request body into the provided struct and limits
// the request size via an io.LimitReader using the maxMemory param.
//
// The Content-Type e.g. "application/xml" and http method are not checked.
//
// If the maxMemory is equal to or less than 0, it's DefaultMaxBodySize.
func DecodeXML(r *http.Request, maxMemory int64, v interface{}) (err error) {
	return decode(r, maxMemory, v, func(reader io.Reader, v interface{}) error {
		return xml.NewDecoder(reader).Decode(v)
	})
}

func filteFlag(s string) string {
	for i, c := range s {
		if c == ' ' || c == ';' {
			return s[:i]
		}
	}
	return s
}

// SetHeader sets the response header.
func SetHeader(w http.ResponseWriter, key, value string) {
	w.Header().Set(key, value)
}

// GetHeader returns the request header.
func GetHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}

// SetContentType is equal to SetContentTypes(w, []string{value}).
func SetContentType(w http.ResponseWriter, value string) {
	w.Header().Set(ContentType, value)
}

// GetContentType returns the Content-Type of the request body.
func GetContentType(r *http.Request) string {
	return filteFlag(r.Header.Get("Content-Type"))
}

// IsWebSocket returns true if the request is websocket. Or returns false.
func IsWebSocket(r *http.Request) bool {
	connection := strings.ToLower(GetHeader(r, "Connection"))
	upgrade := strings.ToLower(GetHeader(r, "Upgrade"))
	if strings.Contains(connection, "upgrade") && upgrade == "websocket" {
		return true
	}
	return false
}

// IsAjax returns true if the request is a AJAX request.
func IsAjax(r *http.Request) bool {
	return GetHeader(r, XRequestedWith) == "XMLHttpRequest"
}

// IsMethod checks whether the request method is the given method.
func IsMethod(r *http.Request, method string) bool {
	return r.Method == strings.ToUpper(method)
}

// GetQuerys returns the query values by the key.
//
// If the key does not exist, return nil.
func GetQuerys(values url.Values, key string) []string {
	return values[key]
}

// GetQuery return the first query value by the key.
//
// If the key does not exist, return "".
func GetQuery(values url.Values, key string) string {
	if vs := GetQuerys(values, key); len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// GetQueryDefault is the same as GetQuery, but return the default if the key
// does not exist.
func GetQueryDefault(values url.Values, key, _default string) string {
	if v := GetQuery(values, key); v != "" {
		return v
	}
	return _default
}

// GetQueryInt gets the first query value and converts it to int.
//
// If the key does not exist. return 0, not an error.
func GetQueryInt(values url.Values, key string) (int, error) {
	if v := GetQuery(values, key); v != "" {
		return types.ToInt(v)
	}
	return 0, nil
}

// GetQueryIntDefault is the same as GetQueryInt, but return the default
// if the key does not exist.
func GetQueryIntDefault(values url.Values, key string, _default int) (int, error) {
	if v, err := GetQueryInt(values, key); err != nil {
		return 0, err
	} else if v != 0 {
		return v, nil
	}
	return _default, nil
}

// GetQueryInt64 gets the first query value and converts it to int64.
//
// If the key does not exist. return 0, not an error.
func GetQueryInt64(values url.Values, key string) (int64, error) {
	if v := GetQuery(values, key); v != "" {
		return types.ToInt64(v)
	}
	return 0, nil
}

// GetQueryInt64Default is the same as GetQueryInt64, but return the default
// if the key does not exist.
func GetQueryInt64Default(values url.Values, key string, _default int64) (int64, error) {
	if v, err := GetQueryInt64(values, key); err != nil {
		return 0, err
	} else if v != 0 {
		return v, nil
	}
	return _default, nil
}

// GetQueryFloat64 gets the first query value and converts it to float64.
//
// If the key does not exist. return 0, not an error.
func GetQueryFloat64(values url.Values, key string) (float64, error) {
	if v := GetQuery(values, key); v != "" {
		return types.ToFloat64(v)
	}
	return 0, nil
}

// GetQueryFloat64Default is the same as GetQueryFloat64, but return the default
// if the key does not exist.
func GetQueryFloat64Default(values url.Values, key string, _default float64) (float64, error) {
	if v, err := GetQueryFloat64(values, key); err != nil {
		return 0, err
	} else if v != 0 {
		return v, nil
	}
	return _default, nil
}

// GetQueryBool gets the first query value and converts it to bool.
//
// For "t", "T", "1", "true", "True", "TRUE", it's true
// For "f", "F", "0", "false", "False", "FALSE", "", it's false.
//
// If the key does not exist. return false, not an error.
func GetQueryBool(values url.Values, key string) (bool, error) {
	if v := GetQuery(values, key); v != "" {
		return types.ToBool(v)
	}
	return false, nil
}

// GetQueryStringSlice returns the query string slice by the key.
//
// The format supports two kinds: (1) more than one key, (2) the value separated
// by the comma. For example:
//
//     // URL: /path/to?key1=v1&key1=v2&key1=v3,v4,v5&key2=value2
//     GetQueryStringSlice(r.URL.Query(), "key1") // => [v1, v2, v3, v4, v5]
//
// If the key does not exist or it value is empty, return an empty string slice.
func GetQueryStringSlice(values url.Values, key string) []string {
	if qs := GetQuerys(values, key); len(qs) > 0 {
		vs := make([]string, 0, len(qs)*2)
		for _, q := range qs {
			for _, v := range strings.Split(q, ",") {
				if v != "" {
					vs = append(vs, v)
				}
			}
		}
	}

	return []string{}
}

// AcceptedLanguages returns an array of accepted languages denoted by
// the Accept-Language header sent by the browser
// NOTE: some stupid browsers send in locales lowercase when all the rest send
// it properly
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

// ClientIP implements a best effort algorithm to return the real client IP,
// it parses X-Real-IP and X-Forwarded-For in order to work properly
// with reverse-proxies such us: nginx or haproxy.
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

// SaveUploadedFile uploads the form file to the specific file dst.
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// GetBody returns the body of the HTTP request or response.
//
// The argument must be http.Request or http.Response, or return an error.
func GetBody(reqOresp interface{}) (data []byte, err error) {
	switch r := reqOresp.(type) {
	case *http.Request:
		return io2.ReadN(r.Body, r.ContentLength)
	case *http.Response:
		defer r.Body.Close()
		return io2.ReadN(r.Body, r.ContentLength)
	default:
		return nil, fmt.Errorf("no *http.Request or *http.Response")
	}
}

// DiscardBody discards the body of the HTTP request or response.
//
// The argument must be http.Request or http.Response, or return an error.
func DiscardBody(reqOresp interface{}) error {
	var body io.ReadCloser
	var length int64

	switch r := reqOresp.(type) {
	case *http.Request:
		body = r.Body
		length = r.ContentLength
	case *http.Response:
		body = r.Body
		length = r.ContentLength
		defer body.Close()
	default:
		return fmt.Errorf("no *http.Request or *http.Response")
	}

	if length < 1 {
		return nil
	}

	_, err := io.CopyN(ioutil.Discard, body, length)
	return err
}

//////////////////////////////////////////////////////////////////////////////
// Send the response

func detectContentType(filename string) (t string) {
	if t = mime.TypeByExtension(filepath.Ext(filename)); t == "" {
		t = OctetStream
	}
	return
}

// Attachment is a helper method for returning an attachement file
// to be downloaded, if you with to open inline see function Inline
func Attachment(w http.ResponseWriter, r io.Reader, filename string) (err error) {
	SetHeader(w, ContentDisposition, "attachment;filename="+filename)
	SetContentType(w, detectContentType(filename))
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, r)
	return
}

// Inline is a helper method for returning a file inline to be rendered/opened
// by the browser.
func Inline(w http.ResponseWriter, r io.Reader, filename string) (err error) {
	SetHeader(w, ContentDisposition, "inline;filename="+filename)
	SetContentType(w, detectContentType(filename))
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, r)
	return
}

// FromReader reads the content from the reader, then renders it to the response.
func FromReader(w http.ResponseWriter, reader io.Reader, status int,
	contentType string) error {
	SetContentType(w, contentType)
	w.WriteHeader(status)
	_, err := io.Copy(w, reader)
	return err
}

// Redirect redirects the request to location.
//
// code must be betwwen 300 and 308, that's [300, 308], or return an error.
func Redirect(w http.ResponseWriter, r *http.Request, code int, location string) error {
	if code < 300 || code > 308 {
		return fmt.Errorf("Cannot redirect with status code %d", code)
	}
	if location == "" {
		location = "/"
	}
	http.Redirect(w, r, location, code)
	return nil
}

// Status writes the response header with the status code.
//
// The returned value is nil forever.
func Status(w http.ResponseWriter, code int) error {
	w.WriteHeader(code)
	return nil
}

// Error renders an error into the response.
//
// If the status is not gived, the default is 500.
func Error(w http.ResponseWriter, err error, code ...int) error {
	if len(code) > 0 {
		return String(w, code[0], "%s", err)
	}
	return String(w, http.StatusInternalServerError, "%s", err)
}

// String renders the format string into the response.
func String(w http.ResponseWriter, status int, format string,
	args ...interface{}) error {
	SetContentType(w, TextPlainCharsetUTF8)
	w.WriteHeader(status)
	_, err := fmt.Fprintf(w, format, args...)
	return err
}

// Bytes renders the content into the response with a Content-Type and code.
func Bytes(w http.ResponseWriter, status int, content []byte,
	contentType ...string) error {
	if len(contentType) > 0 {
		SetContentType(w, contentType[0])
	}
	w.WriteHeader(status)
	_, err := w.Write(content)
	return err
}

// JSON marshals provided interface + returns JSON + status code
func JSON(w http.ResponseWriter, status int, i interface{}) error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return JSONBytes(w, status, data)
}

// JSONBytes returns provided JSON response with status code
func JSONBytes(w http.ResponseWriter, status int, data []byte) (err error) {
	return Bytes(w, status, data, ApplicationJSONCharsetUTF8)
}

// JSONP sends a JSONP response with status code and uses `callback` to
// construct the JSONP payload.
func JSONP(w http.ResponseWriter, status int, i interface{}, callback string) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	SetContentType(w, ApplicationJavaScriptCharsetUTF8)
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
	data, err := xml.Marshal(i)
	if err != nil {
		return err
	}

	return XMLBytes(w, status, data)
}

// XMLBytes returns provided XML response with status code
func XMLBytes(w http.ResponseWriter, status int, data []byte) (err error) {
	SetContentType(w, ApplicationXMLCharsetUTF8)
	w.WriteHeader(status)

	if _, err = w.Write(xmlHeaderBytes); err == nil {
		_, err = w.Write(data)
	}

	return
}
