//  Copyright 2020 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package trace

import (
	"log"
	"strconv"
	"time"
)

var (
	timestampInitial = int64(1514736000000)             //开始时间截 (2018-01-01)
	nodeIdBits       = uint(10)                         //机器id所占的位数
	sequenceBits     = uint(12)                         //序列所占的位数
	nodeIdMax        = int64(-1 ^ (-1 << nodeIdBits))   //支持的最大机器id数量
	sequenceMask     = int64(-1 ^ (-1 << sequenceBits)) //
	nodeIdShift      = sequenceBits                     //机器id左移位数
	timestampShift   = sequenceBits + nodeIdBits        //时间戳左移位数
)

type Trace struct {
	SpinLock

	nodeId    int64
	timestamp int64
	sequence  int64
}

func New(nodeId int64) *Trace {
	return &Trace{nodeId: nodeId, sequence: 0}
}

func (t *Trace) Int64() int64 {
	t.Lock()
	defer t.Unlock()

	if t.nodeId&nodeIdMax == 0 {
		log.Printf("[warn] nodeId has out of range[0, %d] , the traceId may be duplicate\n", nodeIdMax)
	}

	timeNow := time.Now().UnixNano() / 1000000 // 纳秒转毫秒
	if t.timestamp != timeNow {
		t.sequence = 0
	}

	t.sequence++
	if t.sequence&sequenceMask == 0 {
		log.Printf("[warn] sequence has out of range[0, %d] , the traceId may be duplicate\n", sequenceMask)
		t.sequence = 0
	}

	t.timestamp = timeNow
	return (timeNow-timestampInitial)<<timestampShift | t.nodeId<<nodeIdShift | t.sequence
}

func (t *Trace) String() string {
	return strconv.FormatInt(t.Int64(), 16)
}
