package aklapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newLRUFromMap[K comparable, V any](m map[K]V) *lruCache[K, V] {
	ret := newLRUCache[K, V](defCacheSz)
	for k, v := range m {
		ret.Add(k, v)
	}
	return ret
}

func Test_addrResponseCache_Lookup(t *testing.T) {
	defer func() {
		NoCache = false
	}()
	type args struct {
		searchText string
	}
	tests := []struct {
		name     string
		NoCache  bool // if true, set NoCache to true before running test
		c        *lruCache[string, *AddrResponse]
		args     args
		wantResp *AddrResponse
		wantOk   bool
	}{
		{"not in cache",
			false,
			newLRUFromMap(map[string]*AddrResponse{
				"xxx": &AddrResponse{Items: []Address{*testAddr}},
			}),
			args{"yyy"},
			nil,
			false,
		},
		{"cached",
			false,
			newLRUFromMap(map[string]*AddrResponse{
				testAddr.Address: &AddrResponse{Items: []Address{*testAddr}},
			}),
			args{testAddr.Address},
			&AddrResponse{Items: []Address{*testAddr}},
			true,
		},
		{"cached, no cache mode",
			true,
			newLRUFromMap(map[string]*AddrResponse{
				testAddr.Address: &AddrResponse{Items: []Address{*testAddr}},
			}),
			args{testAddr.Address},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NoCache = tt.NoCache
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
		ar         *AddrResponse
	}
	tests := []struct {
		name string
		c    *lruCache[string, *AddrResponse]
		args args
		want *AddrResponse
	}{
		{"add",
			newLRUCache[string, *AddrResponse](10),
			args{testAddr.Address, &AddrResponse{Items: []Address{*testAddr}}},
			&AddrResponse{Items: []Address{*testAddr}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Add(tt.args.searchText, tt.args.ar)

			got, ok := tt.c.Lookup(tt.args.searchText)
			assert.True(t, ok)
			assert.Equal(t, got, tt.want)
		})
	}
}

var (
	tomorrow = &CollectionDayDetailResult{
		Collections: []RubbishCollection{
			{Date: time.Now().Add(24 * time.Hour), Rubbish: true},
			{Date: time.Now().Add(7 * 24 * time.Hour), Rubbish: true},
		},
		Address: testAddr,
	}
	today = &CollectionDayDetailResult{
		Collections: []RubbishCollection{
			{Date: time.Now().Add(1 * time.Minute), Rubbish: true},
			{Date: time.Now().Add(7 * 24 * time.Hour), Rubbish: true},
		},
		Address: testAddr,
	}
	yesterday = &CollectionDayDetailResult{
		Collections: []RubbishCollection{
			{Date: time.Now().Add(-24 * time.Hour), Rubbish: true},
			{Date: time.Now().Add(5 * 24 * time.Hour), Rubbish: true},
		},
		Address: testAddr,
	}
	secondYesterday = &CollectionDayDetailResult{
		Collections: []RubbishCollection{
			{Date: time.Now().Add(24 * time.Hour), Rubbish: true},
			{Date: time.Now().Add(-24 * time.Hour), Rubbish: true},
		},
		Address: testAddr,
	}
)

func newRubbishCacheFromMap(m map[string]*CollectionDayDetailResult) *rubbishResultCache {
	var res = rubbishResultCache{
		lc: newLRUCache[string, *CollectionDayDetailResult](defCacheSz),
	}
	for k, v := range m {
		res.Add(k, v)
	}
	return &res
}

func Test_rubbishResultCache_Lookup(t *testing.T) {
	defer func() {
		NoCache = false
	}()
	type args struct {
		searchText string
	}
	tests := []struct {
		name       string
		NoCache    bool // if true, set NoCache to true before running test
		c          *rubbishResultCache
		args       args
		wantResult *CollectionDayDetailResult
		wantOk     bool
	}{
		{"not in cache (empty)", false, newRubbishCacheFromMap(map[string]*CollectionDayDetailResult{}), args{"blah"}, nil, false},
		{"not in cache",
			false,
			newRubbishCacheFromMap(map[string]*CollectionDayDetailResult{
				testAddr.Address: tomorrow,
			}),
			args{"blah"},
			nil, false},
		{"in cache (future)",
			false,
			newRubbishCacheFromMap(map[string]*CollectionDayDetailResult{
				testAddr.Address: tomorrow,
			}),
			args{testAddr.Address},
			tomorrow, true},
		{"in cache (today)",
			false,
			newRubbishCacheFromMap(map[string]*CollectionDayDetailResult{
				testAddr.Address: today,
			}),
			args{testAddr.Address},
			today, true},
		{"in cache, expired",
			false,
			newRubbishCacheFromMap(map[string]*CollectionDayDetailResult{
				testAddr.Address: yesterday,
			}),
			args{testAddr.Address},
			nil, false},
		{"in cache, second entry expired",
			false,
			newRubbishCacheFromMap(map[string]*CollectionDayDetailResult{
				testAddr.Address: secondYesterday,
			}),
			args{testAddr.Address},
			nil, false},
		{"no cache",
			true,
			newRubbishCacheFromMap(map[string]*CollectionDayDetailResult{
				testAddr.Address: secondYesterday,
			}),
			args{testAddr.Address},
			nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NoCache = tt.NoCache
			gotResult, gotOk := tt.c.Lookup(tt.args.searchText)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("rubbishResultCache.Lookup() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotOk != tt.wantOk {
				t.Errorf("rubbishResultCache.Lookup() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_rubbishResultCache_Add(t *testing.T) {
	type args struct {
		searchText string
		ar         *CollectionDayDetailResult
	}
	tests := []struct {
		name string
		c    *rubbishResultCache
		args args
		want *CollectionDayDetailResult
	}{
		{"add",
			newRubbishCacheFromMap(map[string]*CollectionDayDetailResult{
				testAddr.Address: tomorrow,
			}),
			args{testAddr.Address, tomorrow},
			tomorrow,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Add(tt.args.searchText, tt.args.ar)

			res, ok := tt.c.Lookup(tt.args.searchText)
			assert.True(t, ok)
			assert.Equal(t, tt.want, res)
		})
	}
}
