package http2

// Customized constants
const (
	DefaultMaxBodySize = 32 << 20 // 32MB
)

// HTTP Constant Terms and Variables
const (

	//-------------
	// Media types
	//-------------
	ApplicationJSON                  = "application/json"
	ApplicationJSONCharsetUTF8       = ApplicationJSON + "; " + CharsetUTF8
	ApplicationJavaScript            = "application/javascript"
	ApplicationJavaScriptCharsetUTF8 = ApplicationJavaScript + "; " + CharsetUTF8
	ApplicationXML                   = "application/xml"
	ApplicationXMLCharsetUTF8        = ApplicationXML + "; " + CharsetUTF8
	ApplicationForm                  = "application/x-www-form-urlencoded"
	ApplicationQueryParams           = ""
	ApplicationProtobuf              = "application/protobuf"
	ApplicationMsgpack               = "application/msgpack"
	TextHTML                         = "text/html"
	TextHTMLCharsetUTF8              = TextHTML + "; " + CharsetUTF8
	TextPlain                        = "text/plain"
	TextPlainCharsetUTF8             = TextPlain + "; " + CharsetUTF8
	MultipartForm                    = "multipart/form-data"
	OctetStream                      = "application/octet-stream"

	//---------
	// Charset
	//---------
	CharsetUTF8 = "charset=utf-8"

	//---------
	// Headers
	//---------
	Accept                        = "Accept"
	AcceptedLanguage              = "Accept-Language"
	AcceptEncoding                = "Accept-Encoding"
	Authorization                 = "Authorization"
	ContentDisposition            = "Content-Disposition"
	ContentEncoding               = "Content-Encoding"
	ContentLength                 = "Content-Length"
	ContentType                   = "Content-Type"
	ContentDescription            = "Content-Description"
	ContentTransferEncoding       = "Content-Transfer-Encoding"
	Cookie                        = "Cookie"
	SetCookie                     = "Set-Cookie"
	IfModifiedSince               = "If-Modified-Since"
	LastModified                  = "Last-Modified"
	Location                      = "Location"
	Referer                       = "Referer"
	UserAgent                     = "User-Agent"
	Upgrade                       = "Upgrade"
	Vary                          = "Vary"
	WWWAuthenticate               = "WWW-Authenticate"
	XForwardedProto               = "X-Forwarded-Proto"
	XHTTPMethodOverride           = "X-HTTP-Method-Override"
	XForwardedFor                 = "X-Forwarded-For"
	XRealIP                       = "X-Real-IP"
	XRequestedWith                = "X-Requested-With"
	Server                        = "Server"
	Origin                        = "Origin"
	AccessControlRequestMethod    = "Access-Control-Request-Method"
	AccessControlRequestHeaders   = "Access-Control-Request-Headers"
	AccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	AccessControlAllowMethods     = "Access-Control-Allow-Methods"
	AccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	AccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	AccessControlMaxAge           = "Access-Control-Max-Age"
	Expires                       = "Expires"
	CacheControl                  = "Cache-Control"
	Pragma                        = "Pragma"
	Allow                         = "Allow"

	// Security
	StrictTransportSecurity = "Strict-Transport-Security"
	XContentTypeOptions     = "X-Content-Type-Options"
	XXSSProtection          = "X-XSS-Protection"
	XFrameOptions           = "X-Frame-Options"
	ContentSecurityPolicy   = "Content-Security-Policy"
	XCSRFToken              = "X-CSRF-Token"

	Gzip = "gzip"

	WildcardParam = "*wildcard"

	blank = ""
)
