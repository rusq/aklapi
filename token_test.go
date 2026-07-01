package aklapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBearerTokenValid(t *testing.T) {
	base := time.Date(2026, 5, 22, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name  string
		token bearerToken
		at    time.Time
		want  bool
	}{
		{"empty token", bearerToken{expires: base.Add(time.Hour)}, base, false},
		{"valid before margin", bearerToken{value: "token", expires: base.Add(time.Minute)}, base, true},
		{"near expiry", bearerToken{value: "token", expires: base.Add(30 * time.Second)}, base, false},
		{"expired", bearerToken{value: "token", expires: base.Add(-time.Second)}, base, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.token.valid(tt.at))
		})
	}
}

func TestBearerTokenCache(t *testing.T) {
	base := time.Date(2026, 5, 22, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name          string
		startToken    bearerToken
		wantRequests  int32
		wantToken     string
		wantExpiresAt time.Time
	}{
		{
			name:          "reuses cached token",
			startToken:    bearerToken{value: "cached", expires: base.Add(time.Hour)},
			wantRequests:  0,
			wantToken:     "cached",
			wantExpiresAt: base.Add(time.Hour),
		},
		{
			name:          "refreshes expired token",
			startToken:    bearerToken{value: "expired", expires: base.Add(-time.Second)},
			wantRequests:  1,
			wantToken:     "fresh",
			wantExpiresAt: base.Add(time.Hour),
		},
		{
			name:          "refreshes token near expiry",
			startToken:    bearerToken{value: "stale", expires: base.Add(10 * time.Second)},
			wantRequests:  1,
			wantToken:     "fresh",
			wantExpiresAt: base.Add(time.Hour),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restoreTokenGlobals(t)
			now = func() time.Time { return base }

			token := jwtWithExp(t, base.Add(time.Hour))
			var requests int32
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				atomic.AddInt32(&requests, 1)
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.Write([]byte(`self.__next_f.push([1,"initialToken\":\"` + token + `\""])`))
			}))
			t.Cleanup(ts.Close)

			addrTokenURI = ts.URL
			addrTokenClient = ts.Client()
			addrTokenCache = &bearerTokenCache{token: tt.startToken}

			got, err := cachedAddrToken(t.Context())
			require.NoError(t, err)
			wantToken := tt.wantToken
			if tt.wantRequests > 0 {
				wantToken = token
			}
			assert.Equal(t, wantToken, got)
			assert.Equal(t, tt.wantRequests, atomic.LoadInt32(&requests))
			assert.Equal(t, tt.wantExpiresAt, addrTokenCache.token.expires)
		})
	}
}

func TestBearerTokenCacheConcurrentRefresh(t *testing.T) {
	restoreTokenGlobals(t)
	base := time.Date(2026, 5, 22, 12, 0, 0, 0, time.UTC)
	now = func() time.Time { return base }

	token := jwtWithExp(t, base.Add(time.Hour))
	var requests int32
	start := make(chan struct{})
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requests, 1)
		<-start
		w.Write([]byte(`self.__next_f.push([1,"initialToken\":\"` + token + `\""])`))
	}))
	t.Cleanup(ts.Close)

	addrTokenURI = ts.URL
	addrTokenClient = ts.Client()
	addrTokenCache = &bearerTokenCache{}

	const goroutines = 20
	errs := make(chan error, goroutines)
	results := make(chan string, goroutines)
	for range goroutines {
		go func() {
			got, err := cachedAddrToken(context.Background())
			errs <- err
			results <- got
		}()
	}

	close(start)
	for range goroutines {
		require.NoError(t, <-errs)
		assert.Equal(t, token, <-results)
	}
	assert.Equal(t, int32(1), atomic.LoadInt32(&requests))
}

func TestFetchAddrToken(t *testing.T) {
	restoreTokenGlobals(t)
	base := time.Date(2026, 5, 22, 12, 0, 0, 0, time.UTC)
	now = func() time.Time { return base }

	tests := []struct {
		name      string
		body      string
		status    int
		wantToken string
		wantExp   time.Time
		wantErr   string
	}{
		{
			name:      "JWT expiry",
			body:      `self.__next_f.push([1,"initialToken\":\"` + jwtWithExp(t, base.Add(time.Hour)) + `\""])`,
			status:    http.StatusOK,
			wantToken: jwtWithExp(t, base.Add(time.Hour)),
			wantExp:   base.Add(time.Hour),
		},
		{
			name:      "opaque token uses fallback TTL",
			body:      `self.__next_f.push([1,"initialToken\":\"opaque-token\""])`,
			status:    http.StatusOK,
			wantToken: "opaque-token",
			wantExp:   base.Add(addrTokenFallbackTTL),
		},
		{
			name:    "non-200",
			body:    `error`,
			status:  http.StatusInternalServerError,
			wantErr: "status 500",
		},
		{
			name:    "missing token",
			body:    `self.__next_f.push([1,"{}"])`,
			status:  http.StatusOK,
			wantErr: "initialToken not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, browserAccept, r.Header.Get("Accept"))
				assert.Equal(t, "none", r.Header.Get("Sec-Fetch-Site"))
				w.WriteHeader(tt.status)
				w.Write([]byte(tt.body))
			}))
			t.Cleanup(ts.Close)

			addrTokenURI = ts.URL
			addrTokenClient = ts.Client()

			got, err := fetchAddrToken(t.Context())
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantToken, got.value)
			assert.Equal(t, tt.wantExp, got.expires)
		})
	}
}

func TestExtractInitialToken(t *testing.T) {
	tests := []struct {
		name    string
		html    string
		want    string
		wantErr string
	}{
		{
			name: "escaped Next Flight payload",
			html: `self.__next_f.push([1,"initialToken\":\"jwt-token\",\"children\":[]"])`,
			want: "jwt-token",
		},
		{
			name: "JSON payload",
			html: `{"initialToken":"json-token"}`,
			want: "json-token",
		},
		{
			name: "multiple chunks",
			html: `self.__next_f.push([1,"{}"]);self.__next_f.push([2,"initialToken\":\"second-token\""])`,
			want: "second-token",
		},
		{
			name:    "missing token",
			html:    `self.__next_f.push([1,"{}"])`,
			wantErr: "initialToken not found",
		},
		{
			name:    "malformed empty token",
			html:    `{"initialToken":""}`,
			wantErr: "initialToken not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractInitialToken(strings.NewReader(tt.html))
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestJWTExpiry(t *testing.T) {
	expires := time.Date(2026, 5, 22, 12, 30, 0, 0, time.UTC)

	tests := []struct {
		name    string
		token   string
		want    time.Time
		wantErr string
	}{
		{"valid", jwtWithExp(t, expires), expires, ""},
		{"not JWT", "opaque", time.Time{}, "token is not a JWT"},
		{"bad base64", "header.not*base64.sig", time.Time{}, "decode JWT payload"},
		{"bad JSON", "header." + base64.RawURLEncoding.EncodeToString([]byte(`not-json`)) + ".sig", time.Time{}, "decode JWT claims"},
		{"missing exp", jwtWithPayload(t, map[string]any{"sub": "test"}), time.Time{}, "JWT exp claim missing"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jwtExpiry(tt.token)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func restoreTokenGlobals(t *testing.T) {
	t.Helper()
	oldURI := addrTokenURI
	oldActionID := addrTokenActionID
	oldClient := addrTokenClient
	oldProvider := addrTokenProvider
	oldCache := addrTokenCache
	oldNow := now
	t.Cleanup(func() {
		addrTokenURI = oldURI
		addrTokenActionID = oldActionID
		addrTokenClient = oldClient
		addrTokenProvider = oldProvider
		addrTokenCache = oldCache
		now = oldNow
	})
}

func jwtWithExp(t *testing.T, expires time.Time) string {
	t.Helper()
	return jwtWithPayload(t, map[string]any{"exp": expires.Unix()})
}

func jwtWithPayload(t *testing.T, payload map[string]any) string {
	t.Helper()
	claims, err := json.Marshal(payload)
	require.NoError(t, err)
	return "header." + base64.RawURLEncoding.EncodeToString(claims) + ".signature"
}
