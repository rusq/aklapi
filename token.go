package aklapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	addrTokenFallbackTTL   = 5 * time.Minute
	addrTokenRefreshMargin = 30 * time.Second
	maxTokenPageSize       = 16 << 20
)

var (
	addrTokenURI = `https://www.aucklandcouncil.govt.nz/en/rubbish-recycling/rubbish-recycling-collections/rubbish-recycling-collection-days.html`
	// Kept replaceable in case Auckland Council moves token acquisition back to
	// a server action. The embedded initialToken path does not currently use it.
	addrTokenActionID = ""
	addrTokenClient   = &http.Client{Timeout: 15 * time.Second, Transport: &browserTransport{wrapped: http.DefaultTransport}}
	addrTokenProvider = cachedAddrToken
	addrTokenCache    = &bearerTokenCache{}
)

var initialTokenPatterns = []*regexp.Regexp{
	regexp.MustCompile(`initialToken\\":\\"([^"\\]+)`),
	regexp.MustCompile(`"initialToken"\s*:\s*"([^"]+)`),
	regexp.MustCompile(`initialToken":"([^"]+)`),
}

type bearerToken struct {
	value   string
	expires time.Time
}

func (t bearerToken) valid(at time.Time) bool {
	return t.value != "" && t.expires.After(at.Add(addrTokenRefreshMargin))
}

type bearerTokenCache struct {
	mu    sync.Mutex
	token bearerToken
}

func cachedAddrToken(ctx context.Context) (string, error) {
	return addrTokenCache.tokenValue(ctx)
}

func (c *bearerTokenCache) tokenValue(ctx context.Context) (string, error) {
	at := now()
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.token.valid(at) {
		return c.token.value, nil
	}

	token, err := fetchAddrToken(ctx)
	if err != nil {
		slog.WarnContext(ctx, "address token refresh failed", "error", err)
		return "", err
	}
	c.token = token
	return token.value, nil
}

func fetchAddrToken(ctx context.Context) (bearerToken, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addrTokenURI, nil)
	if err != nil {
		return bearerToken{}, err
	}
	setBrowserDocumentHeaders(req)

	resp, err := addrTokenClient.Do(req)
	if err != nil {
		return bearerToken{}, fmt.Errorf("fetch address token page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return bearerToken{}, fmt.Errorf("fetch address token page: status %d", resp.StatusCode)
	}

	token, err := extractInitialToken(io.LimitReader(resp.Body, maxTokenPageSize))
	if err != nil {
		return bearerToken{}, fmt.Errorf("parse address token page: %w", err)
	}

	expires, err := jwtExpiry(token)
	if err != nil {
		slog.DebugContext(ctx, "using fallback address token expiry", "error", err)
		expires = now().Add(addrTokenFallbackTTL)
	}
	return bearerToken{value: token, expires: expires}, nil
}

func extractInitialToken(r io.Reader) (string, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	for _, pattern := range initialTokenPatterns {
		matches := pattern.FindSubmatch(body)
		if len(matches) == 2 {
			token := strings.TrimSpace(string(matches[1]))
			if token == "" {
				return "", errors.New("empty initialToken")
			}
			return token, nil
		}
	}
	return "", errors.New("initialToken not found")
}

func jwtExpiry(token string) (time.Time, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return time.Time{}, errors.New("token is not a JWT")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("decode JWT payload: %w", err)
	}

	var claims struct {
		Exp int64 `json:"exp"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return time.Time{}, fmt.Errorf("decode JWT claims: %w", err)
	}
	if claims.Exp <= 0 {
		return time.Time{}, errors.New("JWT exp claim missing")
	}
	return time.Unix(claims.Exp, 0).UTC(), nil
}
