package main

import (
	"fmt"
	"math/rand"
	"time"
)

//left 表示数组左边的小标
//right 表示数组右边的下表
func QuickSort(left int, right int, arr []int) {
	l := left
	r := right
	//pivot 是数组中间的数
	pivot := arr[(left+right)/2]

	for l < r {
		//以中轴为基准 小左大右
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}

		//l >= r表示任务完成 退出
		if l >= r {
			break
		}

		//swap
		arr[l], arr[r] = arr[r], arr[l]

		//优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, arr)
	}
	if right > l {
		QuickSort(l, right, arr)
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
	QuickSort(0, len(arr)-1, arr)
	t2 := time.Now().UnixNano()
	fmt.Println("排序总耗时：", t2-t1, "纳秒")
	//fmt.Println(arr)
}
