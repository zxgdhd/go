package sort

import (
	_ "fmt"
	"math/rand"
	_ "strings"
	"testing"
)

/*
测试快速排序
*/
func TestQsort(t *testing.T) {
	array := make([]int, 2e8)
	for index := 0; index < cap(array); index++ {
		array[index] = rand.Int()
	}
	Qsort(array, 0, len(array)-1)
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

/*
测试使用循环的快速排序
*/
func TestQsortUseQueue(t *testing.T) {
	array := make([]int, 2e9)
	for index := 0; index < cap(array); index++ {
		array[index] = rand.Int()
	}
	QsortUseQueue(array)
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
