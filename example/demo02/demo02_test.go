package demo02

import (
	"testing"
	"time"
	"fmt"
)

var command1 = "aaa"
var command2 = "bbb"

func TestAddDelayQueue(t *testing.T) {
	AddDelayQueue(command1, time.Now().Unix() + 40)
	AddDelayQueue(command2, time.Now().Unix())
}

func TestDelayQueueToQueue(t *testing.T) {
	uniqueId := DelayQueueToQueue(time.Now().Unix() + 1)
	if uniqueId == "" {
		t.Error("error")
		return
	}
	uniqueId = DelayQueueToQueue(time.Now().Unix() + 1)
	if uniqueId != "" {
		t.Error("error")
		return
	}
	uniqueId = DelayQueueToQueue(time.Now().Unix() + 41)
	if uniqueId == "" {
		t.Error("error")
	}
}


func TestGetCommandFromQueue(t *testing.T) {
	str := GetCommandFromQueue()
	fmt.Println(str)
	if str == "" {
		t.Error("命令错误")
	}
	str = GetCommandFromQueue()
	fmt.Println(str)
	if str == "" {
		t.Error("命令错误")
	}

	str = GetCommandFromQueue()
	fmt.Println(str)
	if str != "" {
		t.Error("")
	}
}