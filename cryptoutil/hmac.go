package cryptoutil

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func HmacHex(data, key []byte) string {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(data)
	return hex.EncodeToString(hmacSha1.Sum(nil))
}
