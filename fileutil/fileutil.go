package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

func PathExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

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
