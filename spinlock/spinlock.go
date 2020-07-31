//  Copyright 2020 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package spinlock

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	lock uintptr
}

func (sl *SpinLock) Lock() {
	for !atomic.CompareAndSwapUintptr(&sl.lock, 0, 1) {
		runtime.Gosched()
	}
}

func (sl *SpinLock) Unlock() {
	atomic.StoreUintptr(&sl.lock, 0)
}
