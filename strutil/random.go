package strutil

import (
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
)

var letterLength = len(letterBytes)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomText(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int()%letterLength]
	}

	return B2s(b)
}
