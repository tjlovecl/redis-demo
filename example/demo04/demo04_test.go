package demo04

import (
	"testing"
	"fmt"
	"math"
)

func TestBatchAdd(t *testing.T) {
	BatchAdd(1, 1000)
	BatchAdd(500, 1500)
}


func TestCount(t *testing.T) {
	num := Count()
	fmt.Println("预期数目为1500， 实际取出的数目为", num)
	numerator := int64(math.Abs(float64(num-1500)))
	percent := numerator * 1000 / 1500
	fmt.Println("误差千分比为", percent)
	if percent > 10 {
		t.Error("误差过大")
	}
}
