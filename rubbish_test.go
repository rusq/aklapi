package aklapi

import (
	_ "embed"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//go:generate curl -L https://www.aucklandcouncil.govt.nz/en/rubbish-recycling/rubbish-recycling-collections/rubbish-recycling-collection-days/12342478585.html -o test_assets/500-queen-street.html
//go:generate curl -L https://www.aucklandcouncil.govt.nz/en/rubbish-recycling/rubbish-recycling-collections/rubbish-recycling-collection-days/12341511281.html -o test_assets/1-luanda-drive.html

// Test data, run go:generate to update, then update dates in tests
// accordingly.
var (
	//go:embed "test_assets/1-luanda-drive.html"
	taRsd1LuandaDrive string

	//go:embed "test_assets/500-queen-street.html"
	taCom500QueenStreet string
)

func Test_parse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *CollectionDayDetailResult
		wantErr bool
	}{
		{"1 Luanda Drive, Ranui",
			args{strings.NewReader(taRsd1LuandaDrive)},
			&CollectionDayDetailResult{
				Collections: []RubbishCollection{
					{
						Day:        "Tuesday, 16 December",
						Date:       adjustYear(time.Date(0, 12, 16, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    true,
						Recycle:    false,
						FoodScraps: false,
					},
					{
						Day:        "Tuesday, 16 December",
						Date:       adjustYear(time.Date(0, 12, 16, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    false,
						Recycle:    false,
						FoodScraps: true,
					},
					{
						Day:        "Tuesday, 23 December",
						Date:       adjustYear(time.Date(0, 12, 23, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    false,
						Recycle:    true,
						FoodScraps: false,
					},
				},
				Address: nil,
			},
			false},
		{"500 Queen Street, CBD",
			args{strings.NewReader(taCom500QueenStreet)},
			&CollectionDayDetailResult{
				Collections: []RubbishCollection{
					{
						Day:     "Friday, 12 December",
						Date:    adjustYear(time.Date(0, 12, 12, 0, 0, 0, 0, defaultLoc)),
						Rubbish: true,
						Recycle: false,
					},
					{
						Day:     "Friday, 12 December",
						Date:    adjustYear(time.Date(0, 12, 12, 0, 0, 0, 0, defaultLoc)),
						Rubbish: false,
						Recycle: true,
					},
				},
				Address: nil,
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCollectionDayDetail(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		testSrv *httptest.Server
		args    args
		want    *CollectionDayDetailResult
		wantErr bool
	}{
		{"main branch",
			httptest.NewServer(testMux()),
			args{addr: "xxx"},
			&CollectionDayDetailResult{
				Collections: []RubbishCollection{
					{
						Day:        "Tuesday, 16 December",
						Date:       adjustYear(time.Date(0, 12, 16, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    true,
						Recycle:    false,
						FoodScraps: false,
					},
					{
						Day:        "Tuesday, 16 December",
						Date:       adjustYear(time.Date(0, 12, 16, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    false,
						Recycle:    false,
						FoodScraps: true,
					},
					{
						Day:        "Tuesday, 23 December",
						Date:       adjustYear(time.Date(0, 12, 23, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    false,
						Recycle:    true,
						FoodScraps: false,
					},
				},
				Address: &Address{
					ID:      "42",
					Address: "Red Square",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.testSrv.Close()
			oldAddrURI := addrURI
			oldcollectionDayURI := collectionDayURI
			defer func() { addrURI = oldAddrURI; collectionDayURI = oldcollectionDayURI }()
			addrURI = tt.testSrv.URL + "/addr"
			collectionDayURI = tt.testSrv.URL + "/rubbish/?an=%s"
			got, err := CollectionDayDetail(t.Context(), tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CollectionDayDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFetchAndParse_StatusCodeError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("temporary error"))
	}))
	defer srv.Close()

	oldURI := collectionDayURI
	oldClient := collectionHTTPClient
	defer func() {
		collectionDayURI = oldURI
		collectionHTTPClient = oldClient
	}()

	collectionDayURI = srv.URL + "/rubbish/%s"
	collectionHTTPClient = srv.Client()

	_, err := fetchandparse(t.Context(), "42")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCollectionRequestHeaders(t *testing.T) {
	t.Cleanup(func() {
		NoCache = false
	})
	NoCache = true

	t.Run("fetchandparse uses browser headers", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if got := r.Header.Get("User-Agent"); got != userAgent {
				http.Error(w, "missing browser user-agent", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Accept"); got != browserAccept {
				http.Error(w, "missing browser accept", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Accept-Language"); got != browserAcceptLanguage {
				http.Error(w, "missing browser accept-language", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Cache-Control"); got != "max-age=0" {
				http.Error(w, "missing cache-control", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Upgrade-Insecure-Requests"); got != "1" {
				http.Error(w, "missing upgrade-insecure-requests", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Sec-Fetch-Site"); got != "none" {
				http.Error(w, "missing sec-fetch-site", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Sec-Fetch-Mode"); got != "navigate" {
				http.Error(w, "missing sec-fetch-mode", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Sec-Fetch-User"); got != "?1" {
				http.Error(w, "missing sec-fetch-user", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Sec-Fetch-Dest"); got != "document" {
				http.Error(w, "missing sec-fetch-dest", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Sec-CH-UA"); got != secCHUA {
				http.Error(w, "missing sec-ch-ua", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Sec-CH-UA-Mobile"); got != "?0" {
				http.Error(w, "missing sec-ch-ua-mobile", http.StatusNotAcceptable)
				return
			}
			if got := r.Header.Get("Sec-CH-UA-Platform"); got != secCHUAPlatform {
				http.Error(w, "missing sec-ch-ua-platform", http.StatusNotAcceptable)
				return
			}
			w.Write([]byte(taRsd1LuandaDrive))
		}))
		t.Cleanup(srv.Close)

		oldURI := collectionDayURI
		oldClient := collectionHTTPClient
		t.Cleanup(func() {
			collectionDayURI = oldURI
			collectionHTTPClient = oldClient
		})

		collectionDayURI = srv.URL + "/rubbish/%s"
		collectionHTTPClient = &http.Client{
			Timeout:   time.Second,
			Transport: &browserTransport{wrapped: srv.Client().Transport},
		}

		got, err := fetchandparse(t.Context(), "42")
		if err != nil {
			t.Fatalf("fetchandparse() error = %v", err)
		}
		if got == nil {
			t.Fatal("fetchandparse() returned nil result")
		}
	})

	t.Run("collection day detail recovers from 406 with browser headers", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/addr":
				writeAddrJSON(w, AddrResponse{Items: []Address{*testAddr}})
			case "/rubbish/42":
				if r.Header.Get("User-Agent") == userAgent &&
					r.Header.Get("Accept") == browserAccept &&
					r.Header.Get("Accept-Language") == browserAcceptLanguage &&
					r.Header.Get("Cache-Control") == "max-age=0" &&
					r.Header.Get("Upgrade-Insecure-Requests") == "1" &&
					r.Header.Get("Sec-Fetch-Site") == "none" &&
					r.Header.Get("Sec-Fetch-Mode") == "navigate" &&
					r.Header.Get("Sec-Fetch-User") == "?1" &&
					r.Header.Get("Sec-Fetch-Dest") == "document" &&
					r.Header.Get("Sec-CH-UA") == secCHUA &&
					r.Header.Get("Sec-CH-UA-Mobile") == "?0" &&
					r.Header.Get("Sec-CH-UA-Platform") == secCHUAPlatform {
					w.Write([]byte(taRsd1LuandaDrive))
					return
				}
				http.Error(w, "collection API returned status code: 406", http.StatusNotAcceptable)
			default:
				http.NotFound(w, r)
			}
		}))
		t.Cleanup(srv.Close)

		oldAddrURI := addrURI
		oldCollectionDayURI := collectionDayURI
		oldAddrClient := addrHTTPClient
		oldCollectionClient := collectionHTTPClient
		t.Cleanup(func() {
			addrURI = oldAddrURI
			collectionDayURI = oldCollectionDayURI
			addrHTTPClient = oldAddrClient
			collectionHTTPClient = oldCollectionClient
		})

		addrURI = srv.URL + "/addr"
		collectionDayURI = srv.URL + "/rubbish/%s"
		addrHTTPClient = &http.Client{
			Timeout:   time.Second,
			Transport: &browserTransport{wrapped: srv.Client().Transport},
		}
		collectionHTTPClient = &http.Client{
			Timeout:   time.Second,
			Transport: &browserTransport{wrapped: srv.Client().Transport},
		}

		got, err := CollectionDayDetail(t.Context(), "500 Queen Street")
		if err != nil {
			t.Fatalf("CollectionDayDetail() error = %v", err)
		}

		assert.NotNil(t, got)
		assert.Equal(t, testAddr, got.Address)
		assert.Len(t, got.Collections, 3)
	})
}

func TestCollectionDayDetailResult_NextRubbish(t *testing.T) {
	type fields struct {
		Collections []RubbishCollection
		Address     *Address
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{"rubbish",
			fields{
				Collections: []RubbishCollection{
					{Day: "Someday",
						Date:       time.Date(2000, 2, 1, 0, 0, 0, 0, defaultLoc),
						Rubbish:    false,
						Recycle:    true,
						FoodScraps: false,
					},
					{Day: "This is the day",
						Date:       time.Date(1991, 9, 16, 0, 0, 0, 0, defaultLoc),
						Rubbish:    true,
						Recycle:    false,
						FoodScraps: false,
					},
				},
			},
			time.Date(1991, 9, 16, 0, 0, 0, 0, defaultLoc),
		},
		{"not found", fields{Collections: []RubbishCollection{}}, time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &CollectionDayDetailResult{
				Collections: tt.fields.Collections,
				Address:     tt.fields.Address,
			}
			if got := res.NextRubbish(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectionDayDetailResult.NextRubbish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollectionDayDetailResult_NextRecycle(t *testing.T) {
	type fields struct {
		Collections []RubbishCollection
		Address     *Address
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{"rubbish",
			fields{
				Collections: []RubbishCollection{
					{Day: "Someday",
						Date:       time.Date(2000, 2, 1, 0, 0, 0, 0, defaultLoc),
						Rubbish:    true,
						Recycle:    false,
						FoodScraps: false,
					},
					{Day: "This is the day",
						Date:       time.Date(1991, 9, 16, 0, 0, 0, 0, defaultLoc),
						Rubbish:    true,
						Recycle:    true,
						FoodScraps: false,
					},
				},
			},
			time.Date(1991, 9, 16, 0, 0, 0, 0, defaultLoc),
		},
		{"not found", fields{Collections: []RubbishCollection{}}, time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &CollectionDayDetailResult{
				Collections: tt.fields.Collections,
				Address:     tt.fields.Address,
			}
			if got := res.NextRecycle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectionDayDetailResult.NextRecycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRubbishCollection_parseDate(t *testing.T) {
	type fields struct {
		Day     string
		Date    time.Time
		Rubbish bool
		Recycle bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"some date", fields{Day: "Monday, 16 September"}, false},
		{"invalid date", fields{Day: "16 September"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RubbishCollection{
				Day:     tt.fields.Day,
				Date:    tt.fields.Date,
				Rubbish: tt.fields.Rubbish,
				Recycle: tt.fields.Recycle,
			}
			if err := r.parseDate(); (err != nil) != tt.wantErr {
				t.Errorf("RubbishCollection.parseDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func testMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/addr", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(AddrResponse{Items: []Address{*testAddr}})
		if err != nil {
			panic(err)
		}
		w.Write(data)
	})
	mux.HandleFunc("/rubbish/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(taRsd1LuandaDrive))
	})

	return mux
}
