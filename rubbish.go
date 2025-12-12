package aklapi

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// NoCache is a flag to disable caching of the results.
var NoCache = false

const (
	dateLayout = "Monday, 2 January"
)

var (
	// defined as a variable so it can be overridden in tests.
	collectionDayURI = `https://new.aucklandcouncil.govt.nz/en/rubbish-recycling/rubbish-recycling-collections/rubbish-recycling-collection-days/%s.html`
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

func (r *RubbishCollection) String() string {
	return fmt.Sprintf("%s: %s", r.Day, r.Type())
}

func (r *RubbishCollection) Type() string {
	switch {
	case r.Rubbish:
		return "Rubbish"
	case r.Recycle:
		return "Recycle"
	case r.FoodScraps:
		return "Food Scraps"
	default:
		return "Unknown"
	}
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
	result, err := fetchandparse(address.ID)
	if err != nil {
		return nil, err
	}
	result.Address = address
	rubbishCache.Add(addr, result)
	log.Printf("rubbish fetch and parse complete in %s", time.Since(start))
	return result, nil
}

// fetchandparse retrieves the data from the webpage and attempts to parse it.
func fetchandparse(addressID string) (*CollectionDayDetailResult, error) {
	resp, err := http.Get(fmt.Sprintf(collectionDayURI, addressID))
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
	const scheduledCardSelector = ".acpl-schedule-card"
	p.detail = make([]RubbishCollection, 0, 3)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	doc.Find(scheduledCardSelector).Each(func(i int, card *goquery.Selection) {
		// Only parse Household collection card for now, not Commercial
		cardTitle := strings.TrimSpace(card.Find("h4.card-title").Text())
		if !strings.Contains(cardTitle, "Household collection") {
			return
		}
		p.parseScheduleCard(card)
	})

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

// parseScheduleCard parses a schedule card and extracts collection information.
func (p *refuseParser) parseScheduleCard(card *goquery.Selection) {
	// Each collection line is a <p class="mb-0 lead"> containing a span.acpl-icon-with-attribute.left
	card.Find("p.mb-0.lead span.acpl-icon-with-attribute.left").Each(func(i int, span *goquery.Selection) {
		icon := span.Find("i.acpl-icon")
		date := strings.TrimSpace(span.Find("b").First().Text())
		if date == "" { // skip empty (e.g. food scraps absent)
			return
		}
		var rc RubbishCollection
		rc.Day = date
		if icon.HasClass("rubbish") {
			rc.Rubbish = true
		} else if icon.HasClass("recycle") {
			rc.Recycle = true
		} else if icon.HasClass("food-waste") {
			rc.FoodScraps = true
		} else {
			// Unknown icon type; ignore
			return
		}
		p.detail = append(p.detail, rc)
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
