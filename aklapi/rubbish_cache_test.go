package aklapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func Test_rubbishResultCache_Lookup(t *testing.T) {
	type args struct {
		searchText string
	}
	tests := []struct {
		name       string
		c          rubbishResultCache
		args       args
		wantResult *CollectionDayDetailResult
		wantOk     bool
	}{
		{"not in cache (empty)", rubbishResultCache{}, args{"blah"}, nil, false},
		{"not in cache",
			rubbishResultCache{
				testAddr.Address: tomorrow,
			},
			args{"blah"},
			nil, false},
		{"in cache (future)",
			rubbishResultCache{
				testAddr.Address: tomorrow,
			},
			args{testAddr.Address},
			tomorrow, true},
		{"in cache (today)",
			rubbishResultCache{
				testAddr.Address: today,
			},
			args{testAddr.Address},
			today, true},
		{"in cache, expired",
			rubbishResultCache{
				testAddr.Address: yesterday,
			},
			args{testAddr.Address},
			nil, false},
		{"in cache, second entry expired",
			rubbishResultCache{
				testAddr.Address: secondYesterday,
			},
			args{testAddr.Address},
			nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		name                   string
		c                      rubbishResultCache
		args                   args
		wantrubbishResultCache rubbishResultCache
	}{
		{"add",
			rubbishResultCache{
				testAddr.Address: tomorrow,
			},
			args{testAddr.Address, tomorrow},
			rubbishResultCache{
				testAddr.Address: tomorrow,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Add(tt.args.searchText, tt.args.ar)
			assert.Equal(t, tt.wantrubbishResultCache, tt.c)
		})
	}
}
