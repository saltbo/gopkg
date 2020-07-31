package randutil

import (
	"math/rand"
	"time"

	"github.com/saltbo/gopkg/convutil"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
)

var letterLength = len(letterBytes)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int()%letterLength]
	}

	return convutil.B2s(b)
}
