package strutil

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRandString(t *testing.T) {
	assert.Equal(t, len(RandomText(10)), 10)
	assert.NotEqual(t, RandomText(1), RandomText(1))
}
