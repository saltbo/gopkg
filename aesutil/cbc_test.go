package aesutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAESCBCCrypto_Decrypt(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		nonce     string
		plain     string
		encrypted string
	}{
		{
			name:      "ok",
			key:       "1",
			nonce:     "ok",
			plain:     "plain text",
			encrypted: "0LkJ9t/7NVcr5aGLhdHeWw==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewAESCBCCrypto(tt.key, tt.nonce)
			got, err := c.Decrypt(tt.encrypted)
			assert.NoError(t, err)
			assert.Equal(t, tt.plain, got)
		})
	}
}

func TestAESCBCCrypto_Encrypt(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		nonce     string
		plain     string
		encrypted string
	}{
		{
			name:      "ok",
			key:       "1",
			nonce:     "ok",
			plain:     "plain text",
			encrypted: "0LkJ9t/7NVcr5aGLhdHeWw==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewAESCBCCrypto(tt.key, tt.nonce)
			got, err := c.Encrypt(tt.plain)
			assert.NoError(t, err)
			assert.Equal(t, tt.encrypted, got)
		})
	}
}
