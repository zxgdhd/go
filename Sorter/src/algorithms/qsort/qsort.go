package sort

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	_ "time"
)

/*
递归进行快速排序，ｉ３四核，６Ｇ内存对2e8个随机数排序６５秒左右，2e9　out of memory
*/
func Qsort(array []int, left int, right int) {
	l, r := left+1, right
	flag := true
	for r > l {
		if flag {
			if array[r] > array[left] {
				r--
			} else {
				flag = !flag
			}
		} else {
			if array[l] < array[left] {
				l++
			} else {
				flag = !flag
				array[l], array[r] = array[r], array[l]
			}
		}
	}
	var index int
	if array[left] > array[r] {
		array[left], array[r] = array[r], array[left]
		index = r
	} else {
		array[left], array[r-1] = array[r-1], array[left]
		index = r - 1
	}
	if index-left > 1 {
		Qsort(array, left, index-1)
	}
	if right-index > 1 {
		Qsort(array, index+1, right)
	}
}

type node struct {
	left  int
	right int
	next  *node
}

type linkedNode struct {
	head  *node
	tail  *node
	count int
}

func (list *linkedNode) add(nextNode *node) {
	if list.count == 0 {
		list.head = nextNode
		list.tail = nextNode
		list.count++
	} else {
		list.tail.next = nextNode
		list.tail = nextNode
		list.count++
	}
}

func (list *linkedNode) get() (*node, error) {
	if list.count <= 0 {
		return nil, errors.New("链表为空！")
	} else if list.count == 1 {
		returnNode := list.head
		list.head = nil
		list.tail = nil
		list.count--
		return returnNode, nil
	}
	returnNode := list.head
	list.head = list.head.next
	list.count--
	return returnNode, nil

}

func QsortUseQueue(array []int, left, right int) {
	head := &node{left, right, nil}
	nodes := &linkedNode{head, head, 1}
	for nodes.count != 0 {
		currentnode, err := nodes.get()
		if err != nil {
			return
		}
		l, r := currentnode.left+1, currentnode.right
		flag := true
		for r > l {
			if flag {
				if array[r] > array[currentnode.left] {
					r--
				} else {
					flag = !flag
				}
			} else {
				if array[l] < array[currentnode.left] {
					l++
				} else {
					flag = !flag
					array[l], array[r] = array[r], array[l]
				}
			}
		}
		var index int
		if array[currentnode.left] > array[r] {
			array[currentnode.left], array[r] = array[r], array[currentnode.left]
			index = r
		} else {
			array[currentnode.left], array[r-1] = array[r-1], array[currentnode.left]
			index = r - 1
		}
		if currentnode.right-index > 1 {
			next := &node{index + 1, currentnode.right, nil}
			nodes.add(next)
		}
		if index-currentnode.left > 1 {
			currentnode.right = index - 1
			currentnode.next = nil
			nodes.add(currentnode)
		}
	}
}

func QsortUseAllProcs(array []int) {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	length := len(array)
	pizz := length / num
	nodes := make(chan []int, num)
	wg := new(sync.WaitGroup)
	for index := 0; index < num; index++ {
		right := (index + 1) * pizz
		if index == num-1 {
			go goSort(array, index*pizz, length-1, nodes, wg)
		} else {
			go goSort(array, index*pizz, right-1, nodes, wg)
		}
		wg.Add(1)
	}
	for count := 1; count < num; count++ {
		node1 := <-nodes
		node2 := <-nodes
		go MergeSort(node1, node2, nodes, wg)
		wg.Add(1)
	}
	wg.Wait()
	array = <-nodes
	for _, value := range array {
		fmt.Print(value, "\n")
	}
}

func goSort(array []int, left, right int, result chan<- []int, wg *sync.WaitGroup) {
	Qsort(array, left, right)
	result <- array[left : right+1]
	wg.Done()
}

func MergeSort(node1 []int, node2 []int, merged chan<- []int, wg *sync.WaitGroup) {
	l1 := len(node1)
	l2 := len(node2)
	mergeResult := make([]int, l1+l2)
	i1 := 0
	i2 := 0
	var ii int
	for ii, _ := range mergeResult {
		if i1 < l1 && i2 < l2 {
			if node1[i1] < node2[i2] {
				mergeResult[ii] = node1[i1]
				i1++
			} else {
				mergeResult[ii] = node2[i2]
				i2++
			}
			continue
		}
		break
	}
	if i1 == l1 {
		for ; i2 < l2; i2++ {
			mergeResult[ii] = node2[i2]
			ii++
		}
	}
	if i2 == l2 {
		for ; i1 < l1; i1++ {
			mergeResult[ii] = node1[i1]
			ii++
		}
	}
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("l1=%d\tl2=%d\ti1=%d\ti2=%d\n", l1, l2, i1, i2)
		}
	}()
	merged <- mergeResult
	wg.Done()
}
