package aklapi

import (
	"net/http"
	"time"
)

var (
	defaultLoc, _ = time.LoadLocation("Pacific/Auckland") // Auckland is in NZ.
)

// userAgent is sent with all outgoing HTTP requests. The Auckland Council
// collection page returns 406 unless requests resemble a modern browser.
//
// The Chrome major version in userAgent must match the "Chromium" version in
// secCHUA, and the OS in userAgent must match secCHUAPlatform: the council's
// CDN/WAF flags inconsistent fingerprints (see issue #32).
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/158.0.0.0 Safari/537.36"

const (
	defaultAccept         = "application/json, text/html, */*"
	browserAccept         = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"
	browserAcceptLanguage = "en-NZ,en;q=0.9"
	secCHUA               = "\"Chromium\";v=\"158\", \"Not_A Brand\";v=\"24\""
	secCHUAPlatform       = "\"Windows\""
)

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
		r.Header.Set("Accept", defaultAccept)
	}
	if r.Header.Get("Accept-Language") == "" {
		r.Header.Set("Accept-Language", browserAcceptLanguage)
	}

	wrapped := t.wrapped
	if wrapped == nil {
		wrapped = http.DefaultTransport
	}
	return wrapped.RoundTrip(r)
}

// setBrowserDocumentHeaders applies the navigation-style headers required by
// the collection HTML page.
func setBrowserDocumentHeaders(r *http.Request) {
	r.Header.Set("Accept", browserAccept)
	r.Header.Set("Upgrade-Insecure-Requests", "1")
	r.Header.Set("Sec-Fetch-Site", "none")
	r.Header.Set("Sec-Fetch-Mode", "navigate")
	r.Header.Set("Sec-Fetch-User", "?1")
	r.Header.Set("Sec-Fetch-Dest", "document")
	r.Header.Set("Sec-CH-UA", secCHUA)
	r.Header.Set("Sec-CH-UA-Mobile", "?0")
	r.Header.Set("Sec-CH-UA-Platform", secCHUAPlatform)
}
