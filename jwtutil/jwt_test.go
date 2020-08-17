package jwtutil

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var claims = &jwt.StandardClaims{
	Audience:  "",
	ExpiresAt: 0,
	Id:        "",
	IssuedAt:  0,
	Issuer:    "",
	NotBefore: 0,
	Subject:   "",
}

func TestInit(t *testing.T) {
	Init("abc")
	token, err := Issue(claims)
	assert.NoError(t, err)

	jt, err := Verify(token, &jwt.StandardClaims{})
	assert.NoError(t, err)
	assert.Equal(t, jt.Claims, claims)
}
