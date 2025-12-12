package aklapi

import (
	"encoding/json"
	"errors"
	"log"
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
func AddressLookup(addr string) (*AddrResponse, error) {
	return MatchingPropertyAddresses(&AddrRequest{SearchText: addr, PageSize: 10})
}

// MatchingPropertyAddresses wrapper around the AKL Council API.
func MatchingPropertyAddresses(addrReq *AddrRequest) (*AddrResponse, error) {
	cachedAr, ok := addrCache.Lookup(addrReq.SearchText)
	if ok {
		log.Printf("cached address result: %q", cachedAr)
		return cachedAr, nil
	}

	req, err := http.NewRequest("GET", addrURI, nil)
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
	log.Printf("address call complete in %s", time.Since(start))

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

func oneAddress(addr string) (*Address, error) {
	resp, err := AddressLookup(addr)
	if err != nil {
		return nil, err
	}
	// need exactly one address to continue
	if len(resp.Items) != 1 {
		return nil, errors.New("ambiguous or empty address results")
	}
	return &resp.Items[0], nil
}
