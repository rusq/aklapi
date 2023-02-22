package aklapi

import (
	"regexp"
	"time"
)

var (
	defaultLoc, _ = time.LoadLocation("Pacific/Auckland") // Auckland is in NZ.
	dow           = regexp.MustCompile("Monday|Tuesday|Wednesday|Thursday|Friday|Saturday|Sunday")
)
