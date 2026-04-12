//go:build integration

package aklapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestLiveSite_CollectionDayDetail hits the live Auckland Council website and
// verifies that the HTML scraper still returns usable data. Run with:
//
//	go test -tags integration -run TestLiveSite
//
// Useful as a periodic smoke test to catch council website redesigns early.
func TestLiveSite_CollectionDayDetail(t *testing.T) {
	NoCache = true
	t.Cleanup(func() { NoCache = false })

	// 500 Queen Street, Auckland Central — a stable, well-known address.
	const knownAddressID = "12342478585"

	res, err := fetchandparse(context.Background(), knownAddressID)
	require.NoError(t, err, "fetchandparse failed — council page structure may have changed")
	assert.NotEmpty(t, res.Collections, "expected at least one collection entry")
	for _, c := range res.Collections {
		assert.NotEmpty(t, c.Day, "collection entry has empty Day")
		assert.False(t, c.Date.IsZero(), "collection entry has zero Date")
	}
}
