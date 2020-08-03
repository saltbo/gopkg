package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrInSlice(t *testing.T) {
	s := []string{"test1", "test2", "test3", "test4", "test5"}
	assert.False(t, StrInSlice("test", s))
	assert.True(t, StrInSlice("test3", s))
}

func TestBoolFromStr(t *testing.T) {
	for _, trueString := range TrueStrings {
		assert.True(t, BoolFromStr(trueString, false))
	}

	for _, trueString := range FalseStrings {
		assert.False(t, BoolFromStr(trueString, true))
	}

	assert.True(t, BoolFromStr("ok1", true))
	assert.False(t, BoolFromStr("ok1", false))
}
