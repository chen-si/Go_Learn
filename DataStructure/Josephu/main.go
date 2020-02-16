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

//分析思路
//1、编写一个函数，PlayGame(first *Boy,startNo int,countNum int)
//2、最后我们使用一个算法在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int) {
	//2、空链表单独处理
	if first.next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	//留一个，判断startNo <= 小孩的总数
	//2、需要定义一个辅助指针，帮助我们删除小孩

	tail := first
	//3、让tail指向环形链表的最后一个小孩 //帮助我们删除小孩
	for {
		if tail.next == first { //说明tail到了最后的小孩
			break

		}
		tail = tail.next
	}
	//4、让first移动到startNo
	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	//5、开始数countNum 然后就删除first指向的小孩
	for {
		//开始数countNum-1次
		for i := 1; i <= countNum-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.no)
		//删除first指向的小孩
		first = first.next
		tail.next = first
		//判断如果tail == first 说明圈中只有一个小孩了
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈 \n", first.no)
}
func main() {
	first := AddBoy(500)
	ShowBoy(first)
	PlayGame(first, 20, 31)
}
