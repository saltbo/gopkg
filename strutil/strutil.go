package strutil

import (
	"strings"
)

var (
	TrueStrings  = []string{"1", "t", "true", "on", "y", "yes"}
	FalseStrings = []string{"0", "f", "false", "off", "n", "no"}
)

func BoolFromStr(s string, def bool) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if StrInSlice(s, TrueStrings) {
		return true
	} else if StrInSlice(s, FalseStrings) {
		return false
	} else {
		return def
	}
}

func StrInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
