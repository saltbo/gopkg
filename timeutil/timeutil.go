package timeutil

import (
	"strings"
	"time"
)

const FormatLayout = "2006-01-02 15:04:05"

type Item struct {
	Go  string
	Std string
}

var formats = []Item{
	{Std: "YYYY", Go: "2006"},
	{Std: "YY", Go: "06"},
	{Std: "MMMM", Go: "January"},
	{Std: "MMM", Go: "Jan"},
	{Std: "MM", Go: "01"},
	{Std: "DD", Go: "02"},
	{Std: "HH", Go: "15"},
	{Std: "hh", Go: "03"},
	{Std: "h", Go: "3"},
	{Std: "mm", Go: "04"},
	{Std: "m", Go: "4"},
	{Std: "ss", Go: "05"},
	{Std: "s", Go: "5"},
}

// Format returns a textual representation of the time value formatted according
// to layout, which defines the format by showing how the reference time
// Example Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
//
// layout defined by:
//  1. YYYY = 2006，YY = 06
//  2. MM = 01， MMM = Jan，MMMM = January
//  3. DD = 02，
//  4. DDD = Mon，DDDD = Monday
//  5. HH = 15，hh = 03, h = 3
//  6. mm = 04, m = 4
//  7. ss = 05, m = 5
func Format(time time.Time, layout string) string {
	for _, format := range formats {
		layout = strings.Replace(layout, format.Std, format.Go, 1)
	}

	return time.Format(layout)
}
