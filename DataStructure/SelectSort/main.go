package main

import "fmt"

//编写函数SelectSort 完成排序
func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		key := i
		for j := i + 1; j < len(arr); j++ {
			//找出最小值
			if arr[j] < arr[key] {
				key = j
			}
		}
		//最小值与前交换
		arr[i], arr[key] = arr[key], arr[i]
	}
}

func main() {
	//定义一个数组
	arr := []int{10, 34, 19, 100, 80, 64, 56}
	SelectSort(arr)

	fmt.Println(arr)
}
