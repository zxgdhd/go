package sort

import (
	_ "errors"
	"math/rand"
	"sort"
	"testing"
)

func TestSortLib(t *testing.T) {
	array := make([]int, 1e7)
	for index := range array {
		array[index] = rand.Int()
	}
	sort.Ints(array)
	flag := false
	for index := 0; index < len(array)-1; index++ {
		if array[index+1] < array[index] {
			flag = true
			break
		}
	}
	if flag {
		t.Error("排序失败!")
	}
}
