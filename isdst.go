package isdst

import (
	"time"
)

// IsDST returns true if the time given is in DST, false if not
// DST is defined as when the offset from UTC is increased
// Ref: <https://github.com/golang/go/issues/42102>
func IsDST(t time.Time) bool {
	// t
	_, tOffset := t.Zone()

	// January 1
	janYear := t.Year()
	if t.Month() > 6 {
		janYear = janYear + 1
	}
	jan1Location := time.Date(janYear, 1, 1, 0, 0, 0, 0, t.Location())
	_, janOffset := jan1Location.Zone()

	// July 1
	jul1Location := time.Date(t.Year(), 7, 1, 0, 0, 0, 0, t.Location())
	_, julOffset := jul1Location.Zone()

	if tOffset == janOffset {
		return janOffset > julOffset
	}
	return julOffset > janOffset
}
