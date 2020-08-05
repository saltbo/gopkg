package cryptoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHmacHex(t *testing.T) {
	hex := HmacHex([]byte("123"), []byte("key"))
	assert.Equal(t, "d4a5b6721d75a5ac15ec698818c77fe1f6e40187", hex)
}

func TestMd5Hex(t *testing.T) {
	hex := Md5Hex("123")
	assert.Equal(t, "202cb962ac59075b964b07152d234b70", hex)
}

func TestMd5HexShort(t *testing.T) {
	hex := Md5HexShort("123")
	assert.Equal(t, "ac59075b964b0715", hex)
}
