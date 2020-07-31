//  Copyright 2019 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package safemap

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Map struct {
	sm sync.Map

	count int64
}

func New() *Map {
	return &Map{
		count: 0,
	}
}

func (m *Map) Load(key string) (value interface{}, ok bool) {
	return m.sm.Load(key)
}

func (m *Map) MustLoad(key string) interface{} {
	v, ok := m.Load(key)
	if !ok {
		panic(fmt.Errorf("key %s not exist", key))
	}

	return v
}

func (m *Map) Range(f func(key string, value interface{}) bool) {
	m.sm.Range(func(key, value interface{}) bool {
		return f(key.(string), value)
	})
}

func (m *Map) Store(key string, value interface{}) {
	atomic.AddInt64(&m.count, 1)
	m.sm.Store(key, value)
}

func (m *Map) Delete(key string) {
	if m.count > 0 {
		atomic.AddInt64(&m.count, -1)
	}
	m.sm.Delete(key)
}

func (m *Map) Length() int64 {
	return m.count
}

func (m *Map) Empty() bool {
	return m.count == 0
}
