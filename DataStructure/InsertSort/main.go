package main

import "fmt"

//插入排序
func InsertSort(arr []int) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
}

func main() {
	arr := []int{1, 5, 61, 86, 34, 48, 54, 45, 98, 100}
	InsertSort(arr)
	fmt.Println(arr)
}
