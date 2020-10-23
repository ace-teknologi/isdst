package isdst

import (
	"testing"
	"time"
)

// parsedTime is the struct representing a parsed time value.
type testCase struct {
	Unix                 int64
	Year                 int
	Month                time.Month
	Day                  int
	Hour, Minute, Second int // 15:04:05 is 15, 4, 5.
	Nanosecond           int // Fractional second.
	Location             string
	IsDST                bool // flag a DST
}

var utctests = []testCase{
	{0, 1970, time.January, 1, 0, 0, 0, 0, "", false},
	{1221681866, 2008, time.September, 17, 20, 4, 26, 0, "", false},
	{-1221681866, 1931, time.April, 16, 3, 55, 34, 0, "", false},
	{-11644473600, 1601, time.January, 1, 0, 0, 0, 0, "", false},
	{599529660, 1988, time.December, 31, 0, 1, 0, 0, "", false},
	{978220860, 2000, time.December, 31, 0, 1, 0, 0, "", false},
}

var nanoutctests = []testCase{
	{0, 1970, time.January, 1, 0, 0, 0, 1e8, "", false},
	{1221681866, 2008, time.September, 17, 20, 4, 26, 2e8, "", false},
}

var localtests = []testCase{
	{0, 1969, time.December, 31, 16, 0, 0, 0, "America/Los_Angeles", false},
	{1221681866, 2008, time.September, 17, 13, 4, 26, 0, "America/Los_Angeles", true},
	{1969898400, 2032, time.June, 3, 11, 0, 0, 0, "America/Los_Angeles", true},
	{1962871199, 2032, time.March, 14, 1, 59, 59, 0, "America/Los_Angeles", false},
	{1962871200, 2032, time.March, 14, 3, 0, 0, 0, "America/Los_Angeles", true},
	{1962871201, 2032, time.March, 14, 3, 0, 1, 0, "America/Los_Angeles", true},
	{1983430799, 2032, time.November, 7, 1, 59, 59, 0, "America/Los_Angeles", true},
	{1983430800, 2032, time.November, 7, 1, 0, 0, 0, "America/Los_Angeles", false},
	{1983430801, 2032, time.November, 7, 1, 0, 1, 0, "America/Los_Angeles", false},
	{1901721599, 2030, time.April, 7, 2, 59, 59, 0, "Australia/Sydney", true},
	{1901721600, 2030, time.April, 7, 2, 0, 0, 0, "Australia/Sydney", false},
	{1901721601, 2030, time.April, 7, 2, 0, 1, 0, "Australia/Sydney", false},
	{1917446399, 2030, time.October, 6, 1, 59, 59, 0, "Australia/Sydney", false},
	{1917446400, 2030, time.October, 6, 3, 0, 0, 0, "Australia/Sydney", true},
	{1917446401, 2030, time.October, 6, 3, 0, 1, 0, "Australia/Sydney", true},
}

var nanolocaltests = []testCase{
	{0, 1969, time.December, 31, 16, 0, 0, 1e8, "America/Los_Angeles", false},
	{1221681866, 2008, time.September, 17, 13, 4, 26, 3e8, "America/Los_Angeles", true},
}

func same(t time.Time, tc testCase) bool {
	// Check aggregates.
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	if year != tc.Year || month != tc.Month || day != tc.Day ||
		hour != tc.Hour || min != tc.Minute || sec != tc.Second {
		return false
	}
	return true
}

func TestIsDST_SecondsUTC(t *testing.T) {
	for _, tc := range utctests {
		loc, err := time.LoadLocation(tc.Location)
		if err != nil {
			t.Errorf("Cannot load timezone '%s'", tc.Location)
			continue
		}
		ts := time.Unix(tc.Unix, 0).In(loc)
		if !same(ts, tc) {
			t.Errorf("same(%d):", ts.Unix())
			t.Errorf("  want=%+v", tc)
			t.Errorf("  have=%v", ts.Format(time.RFC3339))
		}

		isDST := IsDST(ts)
		if isDST != tc.IsDST {
			t.Errorf("IsDST(%d):  // %s", ts.Unix(), ts.Format(time.RFC3339))
			t.Errorf("  want=%t", tc.IsDST)
			t.Errorf("  have=%t", isDST)
		} else {
			t.Logf("IsDST(%+v)=%t", ts, isDST)
		}
	}
}

func TestIsDST_NanosecondsUTC(t *testing.T) {
	for _, tc := range nanoutctests {
		loc, err := time.LoadLocation(tc.Location)
		if err != nil {
			t.Errorf("Cannot load timezone '%s'", tc.Location)
			continue
		}

		ts := time.Unix(tc.Unix, 0).In(loc)
		if !same(ts, tc) {
			t.Errorf("same(%d):", ts.Unix())
			t.Errorf("  want=%+v", tc)
			t.Errorf("  have=%v", ts.Format(time.RFC3339))
		}

		isDST := IsDST(ts)

		if isDST != tc.IsDST {
			t.Errorf("IsDST(%d):  // %s", ts.Unix(), ts.Format(time.RFC3339))
			t.Errorf("  want=%t", tc.IsDST)
			t.Errorf("  have=%t", isDST)
		} else {
			t.Logf("IsDST(%+v)=%t", ts, isDST)
		}
	}
}

func TestIsDST_SecondsLocal(t *testing.T) {
	for _, tc := range localtests {
		loc, err := time.LoadLocation(tc.Location)
		if err != nil {
			t.Errorf("Cannot load timezone '%s'", tc.Location)
			continue
		}

		ts := time.Unix(tc.Unix, 0).In(loc)
		if !same(ts, tc) {
			t.Errorf("same(%d):", ts.Unix())
			t.Errorf("  want=%+v", tc)
			t.Errorf("  have=%v", ts.Format(time.RFC3339))
		}

		isDST := IsDST(ts)

		if isDST != tc.IsDST {
			t.Errorf("IsDST(%d):  // %s", ts.Unix(), ts.Format(time.RFC3339))
			t.Errorf("  want=%t", tc.IsDST)
			t.Errorf("  have=%t", isDST)
		} else {
			t.Logf("IsDST(%+v)=%t", ts, isDST)
		}
	}
}

func TestIsDST_NanosecondsLocal(t *testing.T) {
	for _, tc := range nanolocaltests {
		loc, err := time.LoadLocation(tc.Location)
		if err != nil {
			t.Errorf("Cannot load timezone '%s'", tc.Location)
			continue
		}

		ts := time.Unix(tc.Unix, 0).In(loc)
		if !same(ts, tc) {
			t.Errorf("same(%d):", ts.Unix())
			t.Errorf("  want=%+v", tc)
			t.Errorf("  have=%v", ts.Format(time.RFC3339))
		}

		isDST := IsDST(ts)

		if isDST != tc.IsDST {
			t.Errorf("IsDST(%d):  // %s", ts.Unix(), ts.Format(time.RFC3339))
			t.Errorf("  want=%t", tc.IsDST)
			t.Errorf("  have=%t", isDST)
		} else {
			t.Logf("IsDST(%+v)=%t", ts, isDST)
		}
	}
}
