package demo06

import (
	"testing"
	"fmt"
	"time"
)

func TestLimit(t *testing.T) {
	for i := 1; i <= TIMES+1; i++ {
		b1 := i < TIMES+1
		b2 := Limit()
		if b1 != b2 {
			t.Error(fmt.Sprintf("第%d次和期望不符，期望%t, 实际%t", i, b1, b2))
		}
	}
	time.Sleep(1*time.Second)
	if b := Limit(); b {
		t.Error(fmt.Sprintf("第%d次和期望不符，期望%t, 实际%t", 5, false, b))
	}
	time.Sleep((INTERVAL+3)*time.Second)
	if b := Limit(); !b {
		t.Error(fmt.Sprintf("第%d次和期望不符，期望%t, 实际%t", 6, true, b))
	}

}
