//  Copyright 2019 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package signer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSign(t *testing.T) {
	s := Hmac("123456", []byte("abcd"))

	assert.Equal(t, "0757052d9e91f749ae03d87b1a6ca86868889292", s.Hash())
	assert.True(t, s.Verify("0757052d9e91f749ae03d87b1a6ca86868889292"))
}

func TestNewMd5Sign(t *testing.T) {
	s := Md5("123456")

	assert.Equal(t, "ZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2U=", s.Hash())
}
