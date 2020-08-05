package httputil

import (
	"net/http"
)

// Hop-by-hop headers. These are removed when sent to the backend.
// As of RFC 7230, hop-by-hop headers are required to appear in the
// Connection header field. These are the headers defined by the
// obsoleted RFC 2616 (section 13.5.1) and are used for backward
// compatibility.
var hopHeaders = []string{
	"Connection",       // section 6.1
	"Proxy-Connection", // non-standard but still sent by libcurl and rejected by e.g. google
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",      // canonicalized version of "TE"
	"Trailer", // not Trailers per URL above; https://www.rfc-editor.org/errata_search.php?eid=4522
	"Transfer-Encoding",
	"Upgrade",
}

// RemoveHopHeaders
func RemoveHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

// HeaderCopyTo
func HeaderCopyTo(src, dst http.Header) {
	RemoveHopHeaders(src)
	for k, v := range src {
		for _, vv := range v {
			if k == "Set-Cookie" {
				dst.Add(k, vv)
			} else {
				dst.Set(k, vv)
			}
		}
	}
}
