package mailutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMail_Send(t *testing.T) {
	Init(Config{
		Host:     "smtpdm.aliyun.com:25",
		Sender:   "no-reply@saltbo.fun",
		Username: "Moreu",
		Password: "23333",
	})
	err := Send("title", "saltbo@foxmail.com", "body")
	assert.Error(t, err)
}
