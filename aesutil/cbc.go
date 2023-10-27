package aesutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/saltbo/gopkg/strutil"
)

type AESCBCCrypto struct {
	key   []byte
	nonce []byte
}

func NewAESCBCCrypto(key, nonce string) *AESCBCCrypto {
	return &AESCBCCrypto{
		key:   []byte(strutil.Md5Hex(key)),
		nonce: []byte(strutil.Md5Hex(nonce)),
	}
}

func (c *AESCBCCrypto) Encrypt(v string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", fmt.Errorf("crypto serializer: %w", err)
	}

	plainBytes := PKCS5Padding([]byte(v), block.BlockSize())
	cipherBytes := make([]byte, len(plainBytes))
	cipher.NewCBCEncrypter(block, c.nonce[:block.BlockSize()]).CryptBlocks(cipherBytes, plainBytes)
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

func (c *AESCBCCrypto) Decrypt(v string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", fmt.Errorf("crypto serializer: %w", err)
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return v, err
	}

	plainBytes := make([]byte, len(cipherBytes))
	cipher.NewCBCDecrypter(block, c.nonce[:block.BlockSize()]).CryptBlocks(plainBytes, cipherBytes)
	return string(PKCS5UnPadding(plainBytes)), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length {
		return src
	}

	return src[:(length - unpadding)]
}
