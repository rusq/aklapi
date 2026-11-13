package aklapi

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBrowserFingerprintConsistent guards against the User-Agent and client
// hint headers drifting apart, which the council's CDN/WAF flags.
func TestBrowserFingerprintConsistent(t *testing.T) {
	uaVer := regexp.MustCompile(`Chrome/(\d+)\.`).FindStringSubmatch(userAgent)
	require.Len(t, uaVer, 2, "userAgent has no Chrome version")
	chVer := regexp.MustCompile(`"Chromium";v="(\d+)"`).FindStringSubmatch(secCHUA)
	require.Len(t, chVer, 2, "secCHUA has no Chromium version")
	assert.Equal(t, uaVer[1], chVer[1], "Chrome version mismatch between userAgent and secCHUA")

	platforms := map[string]string{
		"Windows NT": `"Windows"`,
		"Macintosh":  `"macOS"`,
		"X11; Linux": `"Linux"`,
	}
	var want string
	for token, platform := range platforms {
		if strings.Contains(userAgent, token) {
			want = platform
			break
		}
	}
	require.NotEmpty(t, want, "unrecognised platform in userAgent")
	assert.Equal(t, want, secCHUAPlatform, "platform mismatch between userAgent and secCHUAPlatform")
}
