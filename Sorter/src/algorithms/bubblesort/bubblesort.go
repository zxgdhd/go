package sort

/*
冒泡排序
这里形参是ｓｌｉｃｅ，而ｓｌｉｃｅ是引用类型，所以此处不需要返回值
*/
func BubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
