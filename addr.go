package aklapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

var (
	// defined as a variable so it can be overridden in tests.
	addrURI = `https://www.aucklandcouncil.govt.nz/_vti_bin/ACWeb/ACservices.svc/GetMatchingPropertyAddresses`
)

// AddrRequest is the address request.
type AddrRequest struct {
	ResultCount     int    `json:"ResultCount"`
	SearchText      string `json:"SearchText"`
	RateKeyRequired bool   `json:"RateKeyRequired"`
}

// AddrResponse is the address response.
type AddrResponse []Address

// AddrSuggestion is the address suggestion.
type Address struct {
	ACRateAccountKey string `json:"ACRateAccountKey"`
	Address          string `json:"Address"`
	Suggestion       string `json:"Suggestion"`
}

func (s Address) String() string {
	return "<" + s.Address + " (" + s.ACRateAccountKey + ")>"
}

// AddressLookup is a convenience function to get addresses.
func AddressLookup(addr string) (AddrResponse, error) {
	return MatchingPropertyAddresses(&AddrRequest{SearchText: addr, RateKeyRequired: false, ResultCount: 10})
}

// MatchingPropertyAddresses wrapper around the AKL Council API.
func MatchingPropertyAddresses(addrReq *AddrRequest) (AddrResponse, error) {
	cachedAr, ok := addrCache.Lookup(addrReq.SearchText)
	if ok {
		log.Printf("cached address result: %q", cachedAr)
		return cachedAr, nil
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(addrReq); err != nil {
		return nil, err
	}

	start := time.Now()
	resp, err := http.Post(addrURI, "application/json; charset=UTF-8", &buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Printf("address call complete in %s", time.Since(start))

	dec := json.NewDecoder(resp.Body)
	ar := AddrResponse{}
	if err := dec.Decode(&ar); err != nil {
		return nil, err
	}
	addrCache.Add(addrReq.SearchText, ar)
	return ar, nil
}

func oneAddress(addr string) (*Address, error) {
	resp, err := AddressLookup(addr)
	if err != nil {
		return nil, err
	}
	// need exactly one address to continue
	if len(resp) != 1 {
		return nil, errors.New("ambiguous or empty address results")
	}
	return &resp[0], nil
}
