package main

import "fmt"

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
	arr := []int{1, -56, 48, 4, 95, -72, 61, -34, 17, 46}
	QuickSort(0, len(arr)-1, arr)
	fmt.Println(arr)
}
