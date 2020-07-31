//  Copyright 2020 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package trace

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrace_Int64(t *testing.T) {
	trace := New(1001)
	traceId := trace.Int64()
	assert.IsType(t, new(int64), &traceId)

	var wg sync.WaitGroup
	var ids sync.Map
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			traceId := trace.String()
			_, ok := ids.Load(traceId)
			assert.True(t, !ok, "traceId conflict!!!")

			ids.Store(traceId, 1)
		}()
	}
	wg.Wait()
}

func TestTrace_String(t *testing.T) {
	traceId := New(1002).String()
	assert.IsType(t, new(string), &traceId)
}
