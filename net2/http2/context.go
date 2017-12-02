package http2

import (
	"net/http"
	"net/url"
	"strings"
)

// Render is a HTTP render interface.
type Render interface {
	// Render only writes the body data into the response, which should not
	// write the status code and has no need to set the Content-Type header.
	Render(http.ResponseWriter) error
}

// Context is a wrapper of http.Request and http.ResponseWriter.
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter

	query url.Values
}

// NewContext returns a new Context.
func NewContext(w http.ResponseWriter, r *http.Request) Context {
	return Context{
		Request: r,
		Writer:  w,
		query:   r.URL.Query(),
	}
}

// IsWebsocket returns true if the request is websocket.
func (c Context) IsWebsocket() bool {
	if strings.Contains(strings.ToLower(c.GetHeader("Connection")), "upgrade") &&
		strings.ToLower(c.GetHeader("Upgrade")) == "websocket" {
		return true
	}
	return false
}

// ClientIP returns the client ip.
func (c Context) ClientIP() string {
	return ClientIP(c.Request)
}

// ContentType returns the Content-Type header of the request.
func (c Context) ContentType() string {
	return GetContentType(c.Request)
}

// GetRawData returns the raw body data.
func (c Context) GetRawData() ([]byte, error) {
	return GetBody(c.Request)
}

//////////////////////////////////////////////////////////////////////////////
// Get the request Cookie and Set the response Cookie

// Cookie returns the named cookie provided in the request.
//
// It will return http.ErrNoCookie if there is not the named cookie.
func (c Context) Cookie(name string) (string, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	return url.QueryUnescape(cookie.Value)
}

// SetCookie adds a Set-Cookie header into the response header.
//
// If the cookie is invalid, it will be dropped silently.
func (c Context) SetCookie(name, value, path, domain string, maxAge int, secure,
	httpOnly bool) {
	if path == "" {
		path = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
}

//////////////////////////////////////////////////////////////////////////////
// URL Query

// GetQuerys returns all query values for the given key.
//
// It will return nil if not the key.
func (c Context) GetQuerys(key string) []string {
	return c.query[key]
}

// GetQuery returns the first query value for the given key.
//
// It will return "" if not the key.
func (c Context) GetQuery(key string) string {
	if vs := c.GetQuerys(key); len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// GetQueryWithDefault is equal to GetQuery, but returns the default if not
// the key.
func (c Context) GetQueryWithDefault(key, _default string) string {
	if v := c.GetQuery(key); v != "" {
		return v
	}
	return _default
}

//////////////////////////////////////////////////////////////////////////////
// Get the request header and Set the response header.

// GetHeader returns the request header by the key.
func (c Context) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

// SetHeader will set the response header if value is not empty,
// Or delete the response header by the key.
//
// Notice: if key is "", ignore it.
func (c Context) SetHeader(key, value string) {
	if key == "" {
		return
	}

	if value == "" {
		c.Writer.Header().Del(key)
	} else {
		c.Writer.Header().Set(key, value)
	}
}

/////////////////////////////////////////////////////////////////////////////
// Render the response

// Status writes the response header with the status code.
func (c Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

// File Sends the file to the client.
func (c Context) File(filepath string) {
	http.ServeFile(c.Writer, c.Request, filepath)
}

// Data writes some data into the repsonse body, with a status code.
func (c Context) Data(code int, contentType string, data []byte) error {
	return Bytes(c.Writer, code, contentType, data)
}

// Render renders the content into the response body, with a status code.
func (c Context) Render(code int, contentType string, r Render) error {
	c.Status(code)
	SetContentType(c.Writer, contentType)
	return r.Render(c.Writer)
}

// String renders the format string into the response body, with a status code.
func (c Context) String(code int, format string, args ...interface{}) error {
	return String(c.Writer, code, format, args...)
}

// XML renders the XML into the response body, with a status code.
func (c Context) XML(code int, v interface{}) error {
	return XML(c.Writer, code, v)
}

// JSON renders the JSON into the response body, with a status code.
func (c Context) JSON(code int, v interface{}) error {
	return JSON(c.Writer, code, v)
}
