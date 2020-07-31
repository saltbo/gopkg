//  Copyright 2019 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package logger

import (
	"bufio"
	"bytes"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type bufferWriter2 struct {
	sync.Mutex

	store chan []byte
}

func (c *bufferWriter2) Write(b []byte) (n int, err error) {
	c.store <- b
	return len(b), nil
}

// async write file
func (c *bufferWriter2) Flush(w io.Writer) error {
	for range c.store {
	}
	return nil
}

var bw1 *bufio.Writer
var bw3 *bufferWriter2

func init() {
	mock := new(bytes.Buffer)
	bw1 = bufio.NewWriterSize(mock, 4096)
	bw3 = &bufferWriter2{store: make(chan []byte, 4096)}
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			_ = bw1.Flush()
			_ = bw3.Flush(mock)
		}
	}()
}

func BenchmarkBufWriter_Write(b *testing.B) {
	data := "abc\n"
	b.ResetTimer()
	b.Run("buf1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = bw1.Write([]byte(data))
		}
	})
	b.ResetTimer()
	b.Run("buf3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = bw3.Write([]byte(data))
		}
	})
	time.Sleep(time.Second * 2)
}

func TestSugar(t *testing.T) {
	defer func() {
		assert.Equal(t, "test panicf", recover())
	}()

	logs := []string{
		"test debug",
		"test info",
		"test warn",
		"test error",
		"test panicf",
	}

	buf := new(bytes.Buffer)
	sugaredLogger = BufLogger(buf).Sugar()
	Debugf(logs[0])
	Infof(logs[1])
	Warnf(logs[2])
	Errorf(logs[3])
	Panicf(logs[4])
	for _, log := range logs {
		assert.Contains(t, log, buf.String())
	}
}

func TestNamed(t *testing.T) {
	assert.NoError(t, Setup(Config{Path: "/tmp", Level: "debug"}))
	time.Sleep(time.Second * 3) // wait flush
	Named("access").Info("", zap.String("abc", "123"))
	Close()
}
