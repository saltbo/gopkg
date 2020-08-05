package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	pt, err := time.Parse(FormatLayout, "2019-08-03 19:09:08")
	assert.NoError(t, err)
	assert.Equal(t, "2019-08-03 19:09:08", Format(pt, "YYYY-MM-DD HH:mm:ss"))
	assert.Equal(t, "19-08-03 19:09:08", Format(pt, "YY-MM-DD HH:mm:ss"))
	assert.Equal(t, "19-08-03 07:09:08", Format(pt, "YY-MM-DD hh:mm:ss"))
	assert.Equal(t, "August 03，2019", Format(pt, "MMMM DD，YYYY"))
	assert.Equal(t, "Aug 03，2019", Format(pt, "MMM DD，YYYY"))
	assert.Equal(t, "7:9:8", Format(pt, "h:m:s"))
}
