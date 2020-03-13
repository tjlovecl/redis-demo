package demo01

import (
	"testing"
	"time"
)

func TestDistributedLock(t *testing.T) {
	testFun := func(isTrueSuccess bool) {
		b, err := DistributedLock()
		if err != nil {
			t.Error(err)
			return
		}
		if (b.(bool) != isTrueSuccess) {
			t.Error("error")
			return
		}
	}

	testFun(true)

	testFun(false)
	testFun(false)
	time.Sleep(time.Duration(getLockTime()) * time.Second + 1)
	testFun(true)
	testFun(false)
	testFun(false)

}

