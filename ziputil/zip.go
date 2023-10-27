package zip

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/zip"
	"github.com/sabhiram/go-gitignore"
)

type PackOption struct {
	IgnoreFile string
}

func Pack(srcFile string, zipFileWriter io.Writer, opt *PackOption) error {
	zw := zip.NewWriter(zipFileWriter)
	defer func() {
		_ = zw.Close()
	}()

	gitIgnore, _ := ignore.CompileIgnoreFile(filepath.Join(srcFile, opt.IgnoreFile))
	walker := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if gitIgnore != nil && gitIgnore.MatchesPath(path) {
			fmt.Printf("file %s ignored \n", path)
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		w, err := zw.Create(strings.TrimPrefix(path, filepath.Clean(srcFile)))
		if err != nil {
			return err
		}

		_, err = io.Copy(w, f)
		return err
	}

	return filepath.Walk(srcFile, walker)
}

func Unpack(zipFile string, dir string) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer func() {
		_ = r.Close()
	}()

	for _, file := range r.File {
		reader, err := file.Open()
		if err != nil {
			return err
		}

		path := filepath.Join(dir, file.Name)
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}

		if _, err := io.Copy(f, reader); err != nil {
			return err
		}
	}
	return nil
}
