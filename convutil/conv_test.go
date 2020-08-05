package convutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS2bAndB2s(t *testing.T) {
	r := "s2b and b2s"
	b := S2b(r)
	s := B2s(b)

	assert.Equal(t, []byte(r), b)
	assert.Equal(t, r, s)
}

func TestNextNumberPowerOf2(t *testing.T) {
	assert.Equal(t, uint64(16), NextNumberPowerOf2(10))
	assert.Equal(t, uint64(32), NextNumberPowerOf2(20))
	assert.Equal(t, uint64(32), NextNumberPowerOf2(30))
	assert.Equal(t, uint64(64), NextNumberPowerOf2(40))
	assert.Equal(t, uint64(128), NextNumberPowerOf2(100))
	assert.Equal(t, uint64(256), NextNumberPowerOf2(200))
	assert.Equal(t, uint64(512), NextNumberPowerOf2(500))
	assert.Equal(t, uint64(1024), NextNumberPowerOf2(1000))
}
