package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	assert.True(t, In(1, []int{1, 2, 3}))
	assert.True(t, In("abc", []string{"aaa", "bbb", "abc"}))
	assert.True(t, In(true, []bool{false, false, true}))

	assert.True(t, !In(4, []int{1, 2, 3}))
	assert.True(t, !In("ddd", []string{"aaa", "bbb", "abc"}))
	assert.True(t, !In(true, []bool{false, false, false}))
}
