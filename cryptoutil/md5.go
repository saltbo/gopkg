package cryptoutil

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hex(str string) string {
	hbs := md5.Sum([]byte(str))
	return hex.EncodeToString(hbs[:])
}

func Md5HexShort(str string) string {
	hbs := md5.Sum([]byte(str))
	return hex.EncodeToString(hbs[4:12])
}
