//  Copyright 2019 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package signer

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

type Signer struct {
	hash string
}

// md5 and base64 encode
func Md5(signStr string) *Signer {
	signMd5 := md5.Sum(bytes.NewBufferString(signStr).Bytes())
	md5Bytes := []byte(hex.EncodeToString(signMd5[:]))
	return &Signer{hash: base64.StdEncoding.EncodeToString(md5Bytes)}
}

// hmacSha1 encode
func Hmac(signStr string, key []byte) *Signer {
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(bytes.NewBufferString(signStr).Bytes())
	return &Signer{hash: hex.EncodeToString(hmacSha1.Sum(nil))}
}

func (s *Signer) Hash() string {
	return s.hash
}

func (s *Signer) Verify(hash string) bool {
	return s.hash == hash
}
