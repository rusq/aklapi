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

//go:generate curl -L https://www.aucklandcouncil.govt.nz/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=12342478585 -o test_assets/500-queen-street.html
//go:generate curl -L https://www.aucklandcouncil.govt.nz/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=12341511281 -o test_assets/1-luanda-drive.html

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
						Day:        "Tuesday 27 August",
						Date:       adjustYear(time.Date(0, 8, 27, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    true,
						Recycle:    false,
						FoodScraps: false,
					},
					{
						Day:        "Tuesday 27 August",
						Date:       adjustYear(time.Date(0, 8, 27, 0, 0, 0, 0, defaultLoc)),
						Rubbish:    false,
						Recycle:    false,
						FoodScraps: true,
					},
					{
						Day:        "Tuesday 3 September",
						Date:       adjustYear(time.Date(0, 9, 3, 0, 0, 0, 0, defaultLoc)),
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
						Day:     "Thursday 22 August",
						Date:    adjustYear(time.Date(0, 8, 22, 0, 0, 0, 0, defaultLoc)),
						Rubbish: true,
						Recycle: false,
					},
					{
						Day:     "Thursday 22 August",
						Date:    adjustYear(time.Date(0, 8, 22, 0, 0, 0, 0, defaultLoc)),
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
		w.Write([]byte(taRsd1LuandaDrive))
	})

	return mux
}
