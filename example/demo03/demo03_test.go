package demo03

import (
	"testing"
	"fmt"
)

func TestLogin(t *testing.T) {
	day1 := "20191201"
	day2 := "20191202"
	uids1 := []int{1,3,5,6,9}
	uids2 := []int{1,2,3,4,5}

	for _, v := range uids1 {
		Login(v, day1)
	}

	for _, v := range uids2 {
		Login(v, day2)
	}
}


func TestIsLogin(t *testing.T) {
	day := "20191201"
	uidMap := map[int]bool {
		1: true,
		2: false,
		12: false,
	}

	for uid, b := range uidMap {
		rb := IsLogin(uid, day)
		if b != rb {
			t.Error(fmt.Sprintf("%d状态错误, 期望：%t, 实际：%t", uid, b, rb))
		}
	}
}


func TestGetDayLoginNum(t *testing.T) {
	dayMap := map[string]int64 {
		"20191201": 5,
		"20191202": 5,
		"20180101": 0,
	}

	for day, num := range dayMap {
		num2 := GetDayLoginNum(day)
		if num != num2 {
			t.Error(fmt.Sprintf("%s:状态错误, 期望：%d, 实际：%d", day, num, num2))
		}
	}
}


func TestGetLoginNumByDays(t *testing.T) {
	var count int64 = 3
	list := []string{
		"20191201", "20191202",
	}

	result := GetLoginNumByDays(list)
	if result != count {
		t.Error(fmt.Sprintf("连续两天数据错误, 期望：%d, 实际：%d", count, result))
	}
}