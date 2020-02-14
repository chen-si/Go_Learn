package main

import "fmt"

type Boy struct {
	no   int  //编号
	next *Boy //指向下一个小孩的指针
}

//编写一个函数构成环形的单项链表
//num :小孩的个数
//*Boy :返回第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("num 的值不对")
		return first
	}

	//循环的构建环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		// 1、因为第一个小孩特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.next = first
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first
		}
	}
	return first
}

func ShowBoy(first *Boy) {
	//如果环形链表为空
	if first.next == nil {
		fmt.Println("链表为空，没有小孩!")
		return
	}

	//创建一个指针，帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号：%d ->", curBoy.no)
		//退出条件
		if curBoy.next == first {
			break
		}
		curBoy = curBoy.next
	}
	fmt.Println()
}

func main() {
	first := AddBoy(5)

	ShowBoy(first)
}
