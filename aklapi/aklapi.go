package aklapi

import (
	"regexp"
	"time"
)

var now = time.Now // time generator

var (
	defaultLoc, _ = time.LoadLocation("Pacific/Auckland") // Auckland is in NZ.
	dow           = regexp.MustCompile("Monday|Tuesday|Wednesday|Thursday|Friday|Saturday|Sunday")
)
