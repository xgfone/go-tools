package http2

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/xgfone/go-tools/log2"
)

// Render is a HTTP render interface.
type Render interface {
	// Render only writes the body data into the response, which should not
	// write the status code and has no need to set the Content-Type header.
	Render(http.ResponseWriter) error
}

// Context is a wrapper of http.Request and http.ResponseWriter.
//
// Notice: the Context struct refers to github.com/henrylee2cn/faygo and
// github.com/gin-gonic/gin.
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter

	query url.Values
}

// ContextHandler converts a context handler to http.Handler.
//
// For example,
//
//     func handler(c Context) error {
//          // ...
//     }
//     http.Handle("/", ContextHandler(handler))
func ContextHandler(f func(Context) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := f(NewContext(w, r)); err != nil {
			log2.ErrorF("Failed to handle %q: %s", r.RequestURI, err)
		}
	})
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

// Host returns a host:port of the this request from the client.
func (c Context) Host() string {
	return c.Request.Host
}

// Method returns the request method.
func (c Context) Method() string {
	return c.Request.Method
}

// Domain returns the domain of the client.
func (c Context) Domain() string {
	return strings.Split(c.Request.Host, ":")[0]
}

// Path returns the path of the request URL.
func (c Context) Path() string {
	return c.Request.URL.Path
}

// Proxy returns all the proxys.
func (c Context) Proxy() []string {
	if ip := c.GetHeader(XForwardedFor); ip != "" {
		return strings.Split(ip, ",")
	}
	return []string{}
}

// IsMethod returns true if the request method is the given method.
func (c Context) IsMethod(method string) bool {
	return c.Method() == method
}

// IsAjax returns true if the request is a AJAX request.
func (c Context) IsAjax() bool {
	return c.GetHeader(XRequestedWith) == "XMLHttpRequest"
}

// UserAgent returns the request header "UserAgent".
func (c Context) UserAgent() string {
	return c.GetHeader(UserAgent)
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

// Redirect redirects the request to location.
//
// code must be betwwen 300 and 308, that's [300, 308], or return an error.
func (c Context) Redirect(code int, location string) error {
	if code < 300 || code > 308 {
		return fmt.Errorf("Cannot redirect with status code %d", code)
	}
	if location == "" {
		location = "/"
	}
	http.Redirect(c.Writer, c.Request, location, code)
	return nil
}

// Error renders the error information to the response body.
//
// if having no second argument, the status code is 500.
func (c Context) Error(err error, code ...int) error {
	status := 500
	if len(code) > 0 {
		status = code[0]
	}
	return c.String(status, "%s", err)
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
