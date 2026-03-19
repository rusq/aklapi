package aklapi

import (
	"net/http"
	"time"
)

var (
	defaultLoc, _ = time.LoadLocation("Pacific/Auckland") // Auckland is in NZ.
)

// userAgent is sent with all outgoing HTTP requests. The Auckland Council
// website CDN (Fastly) returns 406 for requests that identify as Go's default
// http client, so we send a browser-compatible value instead.
const userAgent = "Mozilla/5.0 (compatible; aklapi/1.0)"

// browserTransport is an http.RoundTripper that adds browser-like headers to
// every request before forwarding it to the underlying transport.
type browserTransport struct {
	wrapped http.RoundTripper
}

func (t *browserTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request so we don't mutate the caller's copy.
	r := req.Clone(req.Context())
	if r.Header.Get("User-Agent") == "" {
		r.Header.Set("User-Agent", userAgent)
	}
	if r.Header.Get("Accept") == "" {
		r.Header.Set("Accept", "application/json, text/html, */*")
	}
	return t.wrapped.RoundTrip(r)
}
