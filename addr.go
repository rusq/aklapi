package aklapi

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

var (
	// defined as a variable so it can be overridden in tests.
	addrURI = `https://www.aucklandcouncil.govt.nz/nextapi/property`
)

// AddrRequest is the address request.
type AddrRequest struct {
	PageSize   int
	SearchText string
}

// Address is the address and its unique identifier (rate account key).
type Address struct {
	ID      string `json:"ID"`
	Address string `json:"Address"`
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

	start := time.Now()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	slog.DebugContext(ctx, "address call complete", "duration", time.Since(start))

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("address API returned status code: " + strconv.Itoa(resp.StatusCode))
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
