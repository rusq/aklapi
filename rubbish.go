package aklapi

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// NoCache is a flag to disable caching of the results.
var NoCache = false

const (
	dateLayout = "Monday 2 January"
)

var (
	// defined as a variable so it can be overridden in tests.
	collectionDayURI = `https://www.aucklandcouncil.govt.nz/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=%s`
)

var errSkip = errors.New("skip this date")

// RubbishCollection contains the date and type of collection.
type RubbishCollection struct {
	Day        string
	Date       time.Time
	Rubbish    bool
	Recycle    bool
	FoodScraps bool
}

// CollectionDayDetailResult contains the information about Rubbish and
// Recycling collection.
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

// NextFoodScraps returns the next food scraps collection date.
func (res *CollectionDayDetailResult) NextFoodScraps() time.Time {
	for _, r := range res.Collections {
		if r.FoodScraps {
			return r.Date
		}
	}
	return time.Time{}
}

// CollectionDayDetail returns a collection day details for the specified
// address as reported by the Auckland Council Website.
func CollectionDayDetail(addr string) (*CollectionDayDetailResult, error) {
	if cachedRes, ok := rubbishCache.Lookup(addr); ok {
		log.Printf("cached rubbish result for %q", addr)
		return cachedRes, nil
	}
	address, err := oneAddress(addr)
	if err != nil {
		return nil, err
	}
	start := time.Now()
	result, err := fetchandparse(address.ACRateAccountKey)
	if err != nil {
		return nil, err
	}
	result.Address = address
	rubbishCache.Add(addr, result)
	log.Printf("rubbish fetch and parse complete in %s", time.Since(start))
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
	result, err := p.parse(r)
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

// Parse parses the auckland council rubbish webpage.
func (p *refuseParser) parse(r io.Reader) ([]RubbishCollection, error) {
	const datesSection = "#ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_pnlHouseholdBlock"
	p.detail = make([]RubbishCollection, 2)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	_ = doc.Find(datesSection).
		Children().
		Slice(1, 3).
		Each(p.parseLinks) // p.parseLinks populates p.detail
	for i := range p.detail {
		if err := (&p.detail[i]).parseDate(); err != nil {
			if err == errSkip {
				// if we get errSkip - we must have gone beyond the proper links
				// therefore terminate loop and remove the empty element.
				p.detail = p.detail[:i]
				break
			}
			log.Println(err)
			continue
		}
	}
	return p.detail, p.Err
}

// parseLinks parses the links within selection
func (p *refuseParser) parseLinks(el int, sel *goquery.Selection) {
	sel.Children().Each(func(n int, sel *goquery.Selection) {
		switch n {
		case 0:
			if dow.FindString(sel.Text()) == "" {
				log.Println("unable to detect day of week")
				return
			}
			p.detail[el].Day = sel.Text()
		default:
			if sel.Text() == "Rubbish" {
				p.detail[el].Rubbish = true
			} else if sel.Text() == "Food scraps" {
				p.detail[el].FoodScraps = true
			} else if sel.Text() == "Recycle" {
				p.detail[el].Recycle = true
			} else {
				p.Err = fmt.Errorf("parse error: sel.Text = %q, el = %d, n = %d", sel.Text(), el, n)
			}
		}
	})
}

func (r *RubbishCollection) parseDate() error {
	if r.Day == "" {
		return errSkip
	}
	t, err := time.ParseInLocation(dateLayout, r.Day, defaultLoc)
	if err != nil {
		return err
	}
	r.Date = adjustYear(t)
	return nil
}
