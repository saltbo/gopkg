package httputil

import "net/http"

func GetScheme(req *http.Request) string {
	// Can't use `r.Request.URL.Scheme`
	// See: https://groups.google.com/forum/#!topic/golang-nuts/pMUkBlQBDF0
	if req.TLS != nil {
		return "https"
	}

	if scheme := req.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	if scheme := req.Header.Get("X-Forwarded-Protocol"); scheme != "" {
		return scheme
	}
	if ssl := req.Header.Get("X-Forwarded-Ssl"); ssl == "on" {
		return "https"
	}
	if scheme := req.Header.Get("X-Url-Scheme"); scheme != "" {
		return scheme
	}

	return "http"
}
