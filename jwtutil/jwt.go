package jwtutil

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var defaultJWTUtil *JWTUtil

func Init(secret string) {
	defaultJWTUtil = New(secret)
}

func Issue(claims jwt.Claims) (string, error) {
	return defaultJWTUtil.Issue(claims)
}

func Verify(token string, claims jwt.Claims) (*jwt.Token, error) {
	return defaultJWTUtil.Parse(token, claims)
}

type JWTUtil struct {
	secret string
}

func New(secret string) *JWTUtil {
	return &JWTUtil{secret: secret}
}

func (p *JWTUtil) Issue(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(p.secret))
}

func (p *JWTUtil) Parse(token string, claims jwt.Claims) (*jwt.Token, error) {
	validation := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(p.secret), nil
	}

	return jwt.ParseWithClaims(token, claims, validation)
}
