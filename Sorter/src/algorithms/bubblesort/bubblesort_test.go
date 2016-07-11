package sort

import (
	_ "fmt"
	"math/rand"
	_ "strings"
	"testing"
)

/*
测试冒泡排序
*/
func TestBubbleSort(t *testing.T) {
	array := make([]int, 2e7)
	for index := 0; index < cap(array); index++ {
		array[index] = rand.Int()
	}
	BubbleSort(array)
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
	//	str := strings.Join(array, "\t")
	//fmt.Println("sorted data:", str)
}
