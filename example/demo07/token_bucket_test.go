package demo07

import (
	"testing"
	"sync/atomic"
	"fmt"
	"time"
)

func TestTokenBucket_TryConsume(t *testing.T) {
	tb := new(TokenBucket)
	tb.Init(5, 200)
	var succNum int64 = 0
	var failNum int64 = 0
	fun := func() {
		s,f := tb.TryConsume(2)
		atomic.AddInt64(&succNum, s)
		atomic.AddInt64(&failNum, f)
	}

	go fun()
	go fun()
	go fun()
	time.Sleep(200 * time.Millisecond)
	if succNum != 5 || failNum != 1 {
		t.Error(fmt.Sprintf("succNum:预期为5，实际为%d, failNum预期为1,实际为%d", succNum, failNum))
	}
	succNum = 0
	failNum = 0
	fun()
	if succNum != 1 || failNum != 1 {
		t.Error(fmt.Sprintf("succNum:预期为1，实际为%d, failNum预期为1,实际为%d", succNum, failNum))
	}

}


