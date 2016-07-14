package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQsortUseAllProcs(t *testing.T) {
	array := make([]int, 20)
	for index := 0; index < cap(array); index++ {
		array[index] = rand.Int()
	}
	QsortUseAllProcs(array)
	flag := false
	for index := 0; index < len(array)-1; index++ {
		fmt.Printf("%d\n", array[index])
		if array[index] > array[index+1] {
			goto hasError
		}
	}
hasError:
	t.Error("排序失败！")
}
