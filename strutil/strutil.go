package strutil

import (
	"strings"

	"github.com/saltbo/gopkg/sliceutil"
)

var (
	TrueStrings  = []string{"1", "t", "true", "on", "y", "yes"}
	FalseStrings = []string{"0", "f", "false", "off", "n", "no"}
)

func BoolFromStr(s string, def bool) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if sliceutil.In(s, TrueStrings) {
		return true
	} else if sliceutil.In(s, FalseStrings) {
		return false
	} else {
		return def
	}
}

// StrInSlice returns a bool if the value in the slice
// Deprecated: Use sliceutil.In replace it
func StrInSlice(a string, list []string) bool {
	return sliceutil.In(a, list)
}
