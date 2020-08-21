package strutil

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// Md5Hex
func Md5Hex(str string) string {
	hbs := md5.Sum([]byte(str))
	return hex.EncodeToString(hbs[:])
}

// Md5HexShort
func Md5HexShort(str string) string {
	hbs := md5.Sum([]byte(str))
	return hex.EncodeToString(hbs[4:12])
}

// HmacHex
func HmacHex(data, key []byte) string {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(data)
	return hex.EncodeToString(hmacSha1.Sum(nil))
}
