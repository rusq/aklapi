package aklapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testAddr = &Address{
	ID:      "42",
	Address: "Red Square",
}

func TestMatchingPropertyAddresses(t *testing.T) {
	type args struct {
		addrReq *AddrRequest
	}
	tests := []struct {
		name    string
		testSrv *httptest.Server
		args    args
		want    *AddrResponse
		wantErr bool
	}{
		{"main branch",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				writeAddrJSON(w, AddrResponse{Items: []Address{*testAddr}})
			})),
			args{&AddrRequest{
				PageSize:   1,
				SearchText: "red sq",
			}},
			&AddrResponse{Items: []Address{*testAddr}},
			false,
		},
		{"non-200 status code",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				writeAddrJSON(w, AddrResponse{})
			})),
			args{&AddrRequest{}},
			nil,
			true,
		},
		{"post error",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.NotFound(w, r)
			})),
			args{&AddrRequest{}},
			nil,
			true,
		},
		{"json error",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("not a json"))
			})),
			args{&AddrRequest{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetAddressTestGlobals(t)
			t.Cleanup(tt.testSrv.Close)
			addrURI = tt.testSrv.URL
			addrTokenProvider = func(context.Context) (string, error) {
				return "test-token", nil
			}
			got, err := MatchingPropertyAddresses(t.Context(), tt.args.addrReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchingPropertyAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchingPropertyAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		testSrv *httptest.Server
		args    args
		want    *AddrResponse
		wantErr bool
	}{
		{"main branch",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				writeAddrJSON(w, AddrResponse{Items: []Address{*testAddr}})
			})),
			args{"red square"},
			&AddrResponse{Items: []Address{*testAddr}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetAddressTestGlobals(t)
			t.Cleanup(tt.testSrv.Close)
			addrURI = tt.testSrv.URL
			addrTokenProvider = func(context.Context) (string, error) {
				return "test-token", nil
			}
			got, err := AddressLookup(t.Context(), tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Address() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oneAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		testSrv *httptest.Server
		args    args
		want    *Address
		wantErr bool
	}{
		{"one address",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				writeAddrJSON(w, AddrResponse{Items: []Address{*testAddr}})
			})),
			args{"red square"},
			testAddr,
			false,
		},
		{"several address",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				writeAddrJSON(w, AddrResponse{Items: []Address{*testAddr, *testAddr}})
			})),
			args{"red squarex"},
			nil,
			true,
		},
		{"no address",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("{\"items\":[]}"))
			})),
			args{"red squarec"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetAddressTestGlobals(t)
			t.Cleanup(tt.testSrv.Close)
			addrURI = tt.testSrv.URL
			addrTokenProvider = func(context.Context) (string, error) {
				return "test-token", nil
			}
			got, err := oneAddress(t.Context(), tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("oneAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oneAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchingPropertyAddressesAuthorization(t *testing.T) {
	resetAddressTestGlobals(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.Equal(t, "red sq", r.URL.Query().Get("query"))
		assert.Equal(t, "1", r.URL.Query().Get("pageSize"))
		writeAddrJSON(w, AddrResponse{Items: []Address{*testAddr}})
	}))
	t.Cleanup(ts.Close)

	addrURI = ts.URL
	addrTokenProvider = func(context.Context) (string, error) {
		return "test-token", nil
	}

	got, err := MatchingPropertyAddresses(t.Context(), &AddrRequest{PageSize: 1, SearchText: "red sq"})
	require.NoError(t, err)
	assert.Equal(t, &AddrResponse{Items: []Address{*testAddr}}, got)
}

func TestMatchingPropertyAddressesTokenFailure(t *testing.T) {
	resetAddressTestGlobals(t)

	var called bool
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(ts.Close)

	addrURI = ts.URL
	addrTokenProvider = func(context.Context) (string, error) {
		return "", errors.New("token unavailable")
	}

	got, err := MatchingPropertyAddresses(t.Context(), &AddrRequest{SearchText: "red sq"})
	require.Error(t, err)
	assert.Nil(t, got)
	assert.False(t, called)
	assert.Contains(t, err.Error(), "get address API token: token unavailable")
}

func TestMatchingPropertyAddressesStatusError(t *testing.T) {
	resetAddressTestGlobals(t)

	tests := []struct {
		name       string
		body       string
		want       string
		forbidText string
	}{
		{
			name: "json error",
			body: `{"error":"Authorisation error. Current IP address logged."}`,
			want: "address API returned status code: 401: Authorisation error. Current IP address logged.",
		},
		{
			name:       "secret-looking body is redacted",
			body:       `{"error":"bad token Bearer header.eyJleHAiOjE3Nzk0NTQ4MDB9.signature"}`,
			want:       "address API returned status code: 401: bad token Bearer <redacted>",
			forbidText: "header.eyJ",
		},
		{
			name:       "large body is bounded",
			body:       strings.Repeat("x", maxAddressErrorBody+100),
			want:       "address API returned status code: 401: " + strings.Repeat("x", maxAddressErrorBody),
			forbidText: strings.Repeat("x", maxAddressErrorBody+1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetAddressTestGlobals(t)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(tt.body))
			}))
			t.Cleanup(ts.Close)

			addrURI = ts.URL
			addrTokenProvider = func(context.Context) (string, error) {
				return "test-token", nil
			}

			_, err := MatchingPropertyAddresses(t.Context(), &AddrRequest{SearchText: "red sq"})
			require.Error(t, err)
			assert.Equal(t, tt.want, err.Error())
			if tt.forbidText != "" {
				assert.NotContains(t, err.Error(), tt.forbidText)
			}
		})
	}
}

func resetAddressTestGlobals(t *testing.T) {
	t.Helper()
	oldURI := addrURI
	oldClient := addrHTTPClient
	oldProvider := addrTokenProvider
	oldCache := addrCache
	oldNoCache := NoCache
	t.Cleanup(func() {
		addrURI = oldURI
		addrHTTPClient = oldClient
		addrTokenProvider = oldProvider
		addrCache = oldCache
		NoCache = oldNoCache
	})
	addrCache = newLRUCache[string, *AddrResponse](defCacheSz)
	NoCache = false
}

func writeAddrJSON(w io.Writer, r AddrResponse) {
	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
