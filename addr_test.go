package aklapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
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
			defer tt.testSrv.Close()
			oldURI := addrURI
			defer func() { addrURI = oldURI }()
			addrURI = tt.testSrv.URL
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
			defer tt.testSrv.Close()
			oldURI := addrURI
			defer func() { addrURI = oldURI }()
			addrURI = tt.testSrv.URL
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
			defer tt.testSrv.Close()
			oldURI := addrURI
			defer func() { addrURI = oldURI }()
			addrURI = tt.testSrv.URL
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

func writeAddrJSON(w io.Writer, r AddrResponse) {
	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
