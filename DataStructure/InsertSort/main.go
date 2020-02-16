package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	rand.Seed(time.Now().Unix())
	arr := make([]int, 2048)
	for i := 0; i < 2048; i++ {
		n := rand.Intn(10000)
		arr = append(arr, n)
	}

	t1 := time.Now().UnixNano()

	InsertSort(arr)
	t2 := time.Now().UnixNano()
	fmt.Println("排序总耗时：", t2-t1, "纳秒")
}
