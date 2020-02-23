package aklapi

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addrResponseCache_Lookup(t *testing.T) {
	type args struct {
		searchText string
	}
	tests := []struct {
		name     string
		c        addrResponseCache
		args     args
		wantResp AddrResponse
		wantOk   bool
	}{
		{"not in cache",
			addrResponseCache{
				"xxx": AddrResponse{*testAddr},
			},
			args{"yyy"},
			nil,
			false,
		},
		{"cached",
			addrResponseCache{
				testAddr.Address: AddrResponse{*testAddr},
			},
			args{testAddr.Address},
			AddrResponse{*testAddr},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, gotOk := tt.c.Lookup(tt.args.searchText)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("addrResponseCache.Lookup() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotOk != tt.wantOk {
				t.Errorf("addrResponseCache.Lookup() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_addrResponseCache_Add(t *testing.T) {
	type args struct {
		searchText string
		ar         AddrResponse
	}
	tests := []struct {
		name                  string
		c                     addrResponseCache
		args                  args
		wantAddrResponseCache addrResponseCache
	}{
		{"add",
			addrResponseCache{},
			args{testAddr.Address, AddrResponse{*testAddr}},
			addrResponseCache{
				testAddr.Address: AddrResponse{*testAddr},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Add(tt.args.searchText, tt.args.ar)
			assert.Equal(t, tt.wantAddrResponseCache, tt.c)
		})
	}
}
