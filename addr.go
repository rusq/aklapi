package aklapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const maxAddressErrorBody = 4 << 10

var (
	// defined as a variable so it can be overridden in tests.
	addrURI = `https://experience.aucklandcouncil.govt.nz/nextapi/property`
	// defined as a variable so tests can replace it.
	addrHTTPClient = &http.Client{Timeout: 15 * time.Second, Transport: &browserTransport{wrapped: http.DefaultTransport}}
)

// AddrRequest is the address request.
type AddrRequest struct {
	PageSize   int
	SearchText string
}

// Address is the address and its unique identifier (rate account key).
type Address struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

// AddrResponse is the address response.
type AddrResponse struct {
	Items []Address `json:"items"`
}

func (s Address) String() string {
	return "<" + s.Address + " (" + s.ID + ")>"
}

// AddressLookup is a convenience function to get addresses.
func AddressLookup(ctx context.Context, addr string) (*AddrResponse, error) {
	return MatchingPropertyAddresses(ctx, &AddrRequest{SearchText: addr, PageSize: 10})
}

// MatchingPropertyAddresses wrapper around the AKL Council API.
func MatchingPropertyAddresses(ctx context.Context, addrReq *AddrRequest) (*AddrResponse, error) {
	cachedAr, ok := addrCache.Lookup(addrReq.SearchText)
	if ok {
		slog.DebugContext(ctx, "found cached address result", "addr", cachedAr)
		return cachedAr, nil
	}

	token, err := addrTokenProvider(ctx)
	if err != nil {
		return nil, fmt.Errorf("get address API token: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addrURI, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("query", addrReq.SearchText)
	if addrReq.PageSize > 0 {
		q.Add("pageSize", strconv.Itoa(addrReq.PageSize))
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", "Bearer "+token)

	start := time.Now()
	resp, err := addrHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	slog.DebugContext(ctx, "address call complete", "duration", time.Since(start))

	if resp.StatusCode != http.StatusOK {
		return nil, addressStatusError(resp)
	}

	dec := json.NewDecoder(resp.Body)
	var apiResp AddrResponse
	if err := dec.Decode(&apiResp); err != nil {
		return nil, err
	}

	addrCache.Add(addrReq.SearchText, &apiResp)
	return &apiResp, nil
}

func oneAddress(ctx context.Context, addr string) (*Address, error) {
	resp, err := AddressLookup(ctx, addr)
	if err != nil {
		return nil, err
	}
	// need exactly one address to continue
	if len(resp.Items) != 1 {
		return nil, errors.New("ambiguous or empty address results")
	}
	return &resp.Items[0], nil
}

func addressStatusError(resp *http.Response) error {
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxAddressErrorBody))
	if err != nil {
		return fmt.Errorf("address API returned status code: %d", resp.StatusCode)
	}

	msg := sanitizeAddressErrorBody(string(body))
	if msg == "" {
		return fmt.Errorf("address API returned status code: %d", resp.StatusCode)
	}
	return fmt.Errorf("address API returned status code: %d: %s", resp.StatusCode, msg)
}

func sanitizeAddressErrorBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.Map(func(r rune) rune {
		if r == '\n' || r == '\r' || r == '\t' {
			return ' '
		}
		if r < ' ' {
			return -1
		}
		return r
	}, body)
	body = strings.Join(strings.Fields(body), " ")
	if body == "" {
		return ""
	}

	const marker = `"error"`
	if strings.Contains(body, marker) {
		var payload struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal([]byte(body), &payload); err == nil && payload.Error != "" {
			body = payload.Error
		}
	}

	body = redactSecretLikeText(body)
	if len(body) > maxAddressErrorBody {
		body = body[:maxAddressErrorBody]
	}
	return body
}

func redactSecretLikeText(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		trimmed := strings.Trim(word, `"'.,;:()[]{}<>`)
		if strings.EqualFold(trimmed, "Bearer") && i+1 < len(words) {
			words[i+1] = "<redacted>"
			continue
		}
		if strings.Count(trimmed, ".") >= 2 && len(trimmed) > 40 {
			words[i] = strings.Replace(word, trimmed, "<redacted>", 1)
		}
	}
	return strings.Join(words, " ")
}
