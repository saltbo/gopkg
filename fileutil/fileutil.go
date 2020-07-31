package fileutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// PathExist
func PathExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

// Visit
func Visit(dir string, suffix string, visitor func(filename string) error) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, suffix) {
			return visitor(path)
		}

		return nil
	})
}

// MD5Hex returns the file md5 hash hex
func MD5Hex(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return ""
	}

	return hex.EncodeToString(md5hash.Sum(nil)[:])
}

// DetectContentType returns the file content-type
func DetectContentType(filepath string) string {
	mimeType := mime.TypeByExtension(path.Ext(filepath))
	if mimeType != "" {
		return mimeType
	}

	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ""
	}

	return http.DetectContentType(fileData)
}
