package sort

import (
	"errors"
	_ "fmt"
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

func QsortUseQueue(array []int) {
	head := &node{0, len(array) - 1, nil}
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
