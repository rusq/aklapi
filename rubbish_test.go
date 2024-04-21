package aklapi

import (
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
		{"ok",
			args{strings.NewReader(testHTML)},
			&CollectionDayDetailResult{
				Collections: []RubbishCollection{
					{Day: "Sunday 21 April",
						Date:       adjustYear(time.Date(0, 04, 21, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    true,
						Recycle:    false,
						FoodScraps: false},
					{Day: "Sunday 21 April",
						Date:       adjustYear(time.Date(0, 04, 21, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    false,
						Recycle:    true,
						FoodScraps: false},
				},
				Address: nil,
			},
			false},
		{"500 queen ok",
			args{strings.NewReader(testHTMLcommercial)},
			&CollectionDayDetailResult{
				Collections: []RubbishCollection{
					{
						Day:     "Sunday 21 April",
						Date:    adjustYear(time.Date(0, 04, 21, 0, 0, 0, 0, defaultLoc)),
						Rubbish: true,
						Recycle: false,
					},
					{
						Day:     "Sunday 21 April",
						Date:    adjustYear(time.Date(0, 04, 21, 0, 0, 0, 0, defaultLoc)),
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
		{"some date", fields{Day: "Monday 16 September"}, false},
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
						Day:        "Sunday 21 April",
						Date:       adjustYear(time.Date(0, 4, 21, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    true,
						Recycle:    false,
						FoodScraps: false,
					},
					{
						Day:        "Sunday 21 April",
						Date:       adjustYear(time.Date(0, 4, 21, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    false,
						Recycle:    true,
						FoodScraps: false,
					},
				},
				Address: &Address{
					ACRateAccountKey: "42",
					Address:          "Red Square",
					Suggestion:       "Red Square",
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
			addrURI = tt.testSrv.URL + "/addr/"
			collectionDayURI = tt.testSrv.URL + "/rubbish/?an=%s"
			got, err := CollectionDayDetail(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CollectionDayDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func testMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/addr/", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(AddrResponse{*testAddr})
		if err != nil {
			panic(err)
		}
		w.Write(data)
	})
	mux.HandleFunc("/rubbish/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testHTML))
	})

	return mux
}
