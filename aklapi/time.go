package aklapi

import "time"

// Some time magic

var now = time.Now // time generator

// sameWeek returns true if dates t1 and t2 are on the same week
func sameWeek(t1 time.Time, t2 time.Time) bool {
	y1, w1 := t1.ISOWeek()
	y2, w2 := t1.ISOWeek()
	if y1 == y2 && w1 == w2 {
		return true
	}
	return false
}

func adjustYear(t time.Time) time.Time {
	today := now()
	thisYear := today.Year()
	t = t.AddDate(thisYear, 0, 0)
	// year correction
	if t.Before(today) && today.Format("20060102") != t.Format("20060102") {
		t = t.AddDate(1, 0, 0)
	}
	return t
}
