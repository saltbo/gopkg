package goutil

import (
	"fmt"
	"runtime"

	"github.com/saltbo/gopkg/logger"
)

// RunWithRetry
func RunWithRetry(runFunc func(idx int) (retry bool, err error), maxCnt int) error {
	var finalErr error
	for idxCnt := 0; idxCnt < maxCnt; idxCnt++ {
		retry, err := runFunc(idxCnt)
		if maxCnt == 1 {
			return err
		}

		finalErr = err
		if !retry {
			break
		}

		logger.Warnf("exec failed: %s, retry[%d]", err, idxCnt+1)
	}

	return finalErr
}

// SafeGo go with recover
func SafeGo(fn func()) {
	go func() {
		defer func() {
			err := recover()
			if err == nil {
				return
			}

			var stack [4096]byte
			n := runtime.Stack(stack[:], false)
			fmt.Printf("panic err: %v\n\n%s\n", err, stack[:n])
		}()

		fn()
	}()
}
