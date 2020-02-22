package aklapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var collectionDayURI = `https://www.aucklandcouncil.govt.nz/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=%s`

var defaultLoc, _ = time.LoadLocation("Pacific/Auckland") // Auckland is in NZ.

// RubbishCollection contains the date and type of collection.
type RubbishCollection struct {
	Day     string
	Date    time.Time
	Rubbish bool
	Recycle bool
}

type CollectionDayDetailResult struct {
	Collections []RubbishCollection
	Address     *Address
}

// NextRubbish returns the next rubbish collection date.
func (res *CollectionDayDetailResult) NextRubbish() time.Time {
	for _, r := range res.Collections {
		if r.Rubbish {
			return r.Date
		}
	}
	return time.Time{}
}

// NextRecycle returns the next recycle collection date.
func (res *CollectionDayDetailResult) NextRecycle() time.Time {
	for _, r := range res.Collections {
		if r.Recycle {
			return r.Date
		}
	}
	return time.Time{}
}

// CollectionDayDetail returns a collection day details for the specified
// address as reported by the Auckland Council Website.
func CollectionDayDetail(addr string) (*CollectionDayDetailResult, error) {
	address, err := oneAddress(addr)
	if err != nil {
		return nil, err
	}
	result, err := fetchandparse(address.ACRateAccountKey)
	if err != nil {
		return nil, err
	}
	result.Address = address
	return result, nil
}

// fetchandparse retrieves the data from the webpage and attempts to parse it.
func fetchandparse(ACRateAccountKey string) (*CollectionDayDetailResult, error) {
	resp, err := http.Get(fmt.Sprintf(collectionDayURI, ACRateAccountKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return parse(resp.Body)
}

// parse parses data from the reader that contains the body of rubbish and
// recycling page.
func parse(r io.Reader) (*CollectionDayDetailResult, error) {
	p := &refuseParser{}
	result, err := p.Parse(r)
	if err != nil {
		return nil, err
	}
	res := CollectionDayDetailResult{
		Collections: result,
	}
	return &res, nil
}

// refuseParser stateful parser for rubbish page
type refuseParser struct {
	detail []RubbishCollection
	Err    error
}

// Parse parses the webpage.
func (p *refuseParser) Parse(r io.Reader) ([]RubbishCollection, error) {
	const datesSection = "#ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_pnlHouseholdBlock"
	p.detail = make([]RubbishCollection, 2)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	_ = doc.Find(datesSection).
		Children().
		Slice(1, 3).
		Each(p.parseLinks)
	for i := range p.detail {
		if err := (&p.detail[i]).parseDate(); err != nil {
			log.Println(err)
			continue
		}
	}
	return p.detail, p.Err
}

func (r *RubbishCollection) parseDate() error {
	datefmt := "Monday 2 January"

	t, err := time.ParseInLocation(datefmt, r.Day, defaultLoc)
	if err != nil {
		return err
	}
	r.Date = adjustYear(t)
	return nil
}

func adjustYear(t time.Time) time.Time {
	thisYear := time.Now().Year()
	t = t.AddDate(thisYear, 0, 0)
	// year correction
	if t.Before(time.Now()) {
		t = t.AddDate(1, 0, 0)
	}
	return t
}

func (p *refuseParser) parseLinks(el int, sel *goquery.Selection) {
	sel.Children().Each(func(n int, sel *goquery.Selection) {
		switch n {
		default:
			p.Err = fmt.Errorf("parse error: sel.Text = %q, el = %d, n = %d", sel.Text(), el, n)
		case 0:
			p.detail[el].Day = sel.Text()
		case 1:
			if sel.Text() != "Rubbish" {
				p.Err = fmt.Errorf("unknown text in rubbish block: %s", sel.Text())
			}
			p.detail[el].Rubbish = true
		case 2:
			if sel.Text() != "Recycle" {
				p.Err = fmt.Errorf("unknown text in rubbish block: %s", sel.Text())
			}
			p.detail[el].Recycle = true
		}
	})
}
