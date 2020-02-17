package main

import (
	"fmt"
)

//使用数组模拟栈
type Stack struct {
	MaxTop int //表示栈的最大可以存放数的个数
	Top    int //表示栈顶
	arr    [5]int
}

func (this *Stack) Push(val int) {
	//先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满！")
		return
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("栈空!")
		return
	}
	for i := 0; i <= this.Top; i++ {
		fmt.Printf(" %d ->", this.arr[i])
	}
	fmt.Println()
}

func (this *Stack) Pop() (val int) {
	if this.Top == -1 {
		fmt.Println("栈空!")
		return
	}
	val = this.arr[this.Top]
	this.Top--
	return
}

func main() {
	stack := &Stack{
		MaxTop: 5,  //表示最多存放5个数据
		Top:    -1, //当栈顶为-1时 表示栈为空
	}
	// stack.Push(1)
	// stack.Push(2)
	// stack.Push(3)
	// stack.Push(4)
	// stack.Push(5)
	// stack.Push(6)

	stack.List()

	fmt.Println("取出了:", stack.Pop())

	stack.List()
}
