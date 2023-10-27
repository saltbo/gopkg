package zip

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPack(t *testing.T) {
	f := md5.New()
	assert.NoError(t, Pack("./testdata/root", f, &PackOption{}))
	assert.Equal(t, "43a9b840a88da2c178de79a7d7e5da89", hex.EncodeToString(f.Sum(nil)))
}

func TestUnpack(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "faasapi-")
	assert.NoError(t, err)
	assert.NoError(t, Unpack("./testdata/test.zip", tmpDir))

	fc, err := os.ReadFile(filepath.Join(tmpDir, "abc.txt"))
	assert.NoError(t, err)
	assert.Equal(t, "this is root file", string(fc))
}
