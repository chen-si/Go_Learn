package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	rand.Seed(time.Now().Unix())
	arr := make([]int, 2048)
	for i := 0; i < 2048; i++ {
		n := rand.Intn(10000)
		arr = append(arr, n)
	}

	t1 := time.Now().UnixNano()

	SelectSort(arr)
	t2 := time.Now().UnixNano()
	fmt.Println("排序总耗时：", t2-t1, "纳秒")
}
