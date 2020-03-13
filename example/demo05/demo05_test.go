package demo05

import (
	"testing"
	"fmt"
)

const KEY = "BLOOM_KEY"

func TestBatchAdd(t *testing.T) {
	values1 := []interface{}{}
	for i := 0; i < 1000; i = i + 2 {
		values1 = append(values1, i)
	}
	values2 := []interface{}{}
	for i := 0; i < 1000; i = i + 3 {
		values2 = append(values2, i)
	}

	BatchAdd(KEY, values1)
	BatchAdd(KEY, values2)
}

func TestBatchExists(t *testing.T) {
	errorList := []int{}
	values := []interface{}{}
	for i := 0; i < 1000; i++ {
		values = append(values, i)
	}

	result := BatchExists(KEY, values)

	for k, v := range result {
		b1 := k%2 == 0 || k%3 == 0
		b2 := v.(int64) == 1
		if b1 != b2 {
			errorList = append(errorList, k)
			fmt.Println(fmt.Sprintf("值：%d, 期望：%t, 实际：%t", k, b1, b2))
		}
	}

	p := len(errorList) * 1000 / len(result)
	fmt.Println(fmt.Sprintf("误差为千分之%d", p))
	if p > 10 {
		t.Error(fmt.Sprintf("误差过大，为%d", p))
	}

}
