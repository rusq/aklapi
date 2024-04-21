package aklapi

import (
	_ "embed"
)

//go:generate curl -o test_assets/your_collection_day_500_queen_street.html https://www.aucklandcouncil.govt.nz/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=12342478585

var (
	// testHTML is the test HTML content for private household.
	//
	//go:embed "test_assets/your_collection_day_500_queen_street.html"
	testHTML           string
	testHTMLcommercial = testHTML
)
