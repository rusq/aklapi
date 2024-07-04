package aklapi

import "time"

// Some time magic

var now = time.Now // time generator

func adjustYear(t time.Time) time.Time {
	today := now()
	thisYear := today.Year()
	t = t.AddDate(thisYear, 0, 0)
	// year correction if date is over a week ago
	if t.Before(today.AddDate(0, 0, -7)) && today.Format("20060102") != t.Format("20060102") {
		t = t.AddDate(1, 0, 0)
	}
	return t
}
