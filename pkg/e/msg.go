package e

var MsgFlags = map[int]string{
	ERROR_DESC_NOT_FOUND:                     "error description not found !",
	HTTP_100_CONTINUE:                        "The HTTP 100 Continue informational status response code indicates that everything so far is OK and that the client should continue with the request or ignore it if it is already finished",
	HTTP_101_SWITCHING_PROTOCOLS:             "The HTTP 101 Switching Protocols response code indicates the protocol the server is switching to as requested by a client which sent the message including the Upgrade request header",
	HTTP_200_OK:                              "The HTTP 200 OK success status response code indicates that the request has succeeded. A 200 response is cacheable by default",
	HTTP_201_CREATED:                         "The HTTP 201 Created success status response code indicates that the request has succeeded and has led to the creation of a resource. The new resource is effectively created before this response is sent back. and the new resource is returned in the body of the message, its location being either the URL of the request, or the content of the Location header",
	HTTP_202_ACCEPTED:                        "The HyperText Transfer Protocol (HTTP) 202 Accepted response status code indicates that the request has been received but not yet acted upon. It is non-committal, meaning that there is no way for the HTTP to later send an asynchronous response indicating the outcome of processing the request. It is intended for cases where another process or server handles the request, or for batch processing",
	HTTP_203_NON_AUTHORITATIVE_INFORMATION:   "The HTTP 203 Non-Authoritative Information response status indicates that the request was successful but the enclosed payload has been modified from that of the origin server's 200 (OK) response by a transforming proxy",
	HTTP_204_NO_CONTENT:                      "The HTTP 204 No Content success status response code indicates that the request has succeeded, but that the client doesn't need to go away from its current page. A 204 response is cacheable by default. An ETag header is included in such a response",
	HTTP_205_RESET_CONTENT:                   "The HTTP 205 Reset Content response status tells the client to reset the document view, so for example to clear the content of a form, reset a canvas state, or to refresh the UI",
	HTTP_206_PARTIAL_CONTENT:                 "The HTTP 206 Partial Content success status response code indicates that the request has succeeded and has the body contains the requested ranges of data, as described in the Range header of the request",
	HTTP_207_MULTI_STATUS:                    "",
	HTTP_300_MULTIPLE_CHOICES:                "The HTTP 300 Multiple Choices redirect status response code indicates that the request has more than one possible responses. The user-agent or the user should choose one of them. As there is no standardized way of choosing one of the responses, this response code is very rarely used",
	HTTP_301_MOVED_PERMANENTLY:               "The HyperText Transfer Protocol (HTTP) 301 Moved Permanently redirect status response code indicates that the resource requested has been definitively moved to the URL given by the Location headers. A browser redirects to this page and search engines update their links to the resource (in 'SEO-speak', it is said that the 'link-juice' is sent to the new URL)",
	HTTP_302_FOUND:                           "The HyperText Transfer Protocol (HTTP) 302 Found redirect status response code indicates that the resource requested has been temporarily moved to the URL given by the Location header",
	HTTP_303_SEE_OTHER:                       "The HyperText Transfer Protocol (HTTP) 303 See Other redirect status response code indicates that the redirects don't link to the newly uploaded resources but to another page, like a confirmation page or an upload progress page. This response code is usually sent back as a result of PUT or POST. The method used to display this redirected page is always GET",
	HTTP_304_NOT_MODIFIED:                    "The HTTP 304 Not Modified client redirection response code indicates that there is no need to retransmit the requested resources. It is an implicit redirection to a cached resource. This happens when the request method is safe, like a GET or a HEAD request, or when the request is conditional and uses a If-None-Match or a If-Modified-Since header",
	HTTP_305_USE_PROXY:                       "",
	HTTP_306_RESERVED:                        "",
	HTTP_307_TEMPORARY_REDIRECT:              "",
	HTTP_400_BAD_REQUEST:                     "The HyperText Transfer Protocol (HTTP) 400 Bad Request response status code indicates that the server could not understand the request due to invalid syntax",
	HTTP_401_UNAUTHORIZED:                    "The HTTP 401 Unauthorized client error status response code indicates that the request has not been applied because it lacks valid authentication credentials for the target resource",
	HTTP_402_PAYMENT_REQUIRED:                "",
	HTTP_403_FORBIDDEN:                       "The HTTP 403 Forbidden client error status response code indicates that the server understood the request but refuses to authorize it",
	HTTP_404_NOT_FOUND:                       "The HTTP 404 Not Found client error response code indicates that the server can't find the requested resource.",
	HTTP_405_METHOD_NOT_ALLOWED:              "The HyperText Transfer Protocol (HTTP) 405 Method Not Allowed response status code indicates that the request method is known by the server but is not supported by the target resource",
	HTTP_406_NOT_ACCEPTABLE:                  "",
	HTTP_407_PROXY_AUTHENTICATION_REQUIRED:   "",
	HTTP_408_REQUEST_TIMEOUT:                 "",
	HTTP_409_CONFLICT:                        "",
	HTTP_410_GONE:                            "",
	HTTP_411_LENGTH_REQUIRED:                 "",
	HTTP_412_PRECONDITION_FAILED:             "",
	HTTP_413_REQUEST_ENTITY_TOO_LARGE:        "",
	HTTP_414_REQUEST_URI_TOO_LONG:            "",
	HTTP_415_UNSUPPORTED_MEDIA_TYPE:          "",
	HTTP_416_REQUESTED_RANGE_NOT_SATISFIABLE: "",
	HTTP_417_EXPECTATION_FAILED:              "",
	HTTP_422_UNPROCESSABLE_ENTITY:            "",
	HTTP_423_LOCKED:                          "",
	HTTP_424_FAILED_DEPENDENCY:               "",
	HTTP_428_PRECONDITION_REQUIRED:           "",
	HTTP_429_TOO_MANY_REQUESTS:               "",
	HTTP_431_REQUEST_HEADER_FIELDS_TOO_LARGE: "",
	HTTP_451_UNAVAILABLE_FOR_LEGAL_REASONS:   "",
	HTTP_500_INTERNAL_SERVER_ERROR:           "",
	HTTP_501_NOT_IMPLEMENTED:                 "",
	HTTP_502_BAD_GATEWAY:                     "",
	HTTP_503_SERVICE_UNAVAILABLE:             "",
	HTTP_504_GATEWAY_TIMEOUT:                 "",
	HTTP_505_HTTP_VERSION_NOT_SUPPORTED:      "",
	HTTP_507_INSUFFICIENT_STORAGE:            "",
	HTTP_511_NETWORK_AUTHENTICATION_REQUIRED: "",

	// //
	ERROR_UNEXPECTED_JSON_INPUT: "unexpected end of JSON input",
	ERROR_NULL_USERNAME:         "username cannot be null when getting the jwt token",
	ERROR_GENERATE_JWT_TOKEN:    "token generate error",
	ERROR_LOGININFO_VALIDATION:  "user not exists or password not correct",
	ERROR_JWT_TOKEN_VALIDATION:  "jwt token verified failed",
	ERROR_GET_JWT_HEADER_FAILED: "get http head Authorization token failed ",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR_DESC_NOT_FOUND]
}
