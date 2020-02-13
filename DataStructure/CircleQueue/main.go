package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxsize int //4
	array   [5]int
	front   int //队首
	rear    int //队尾
}

func (this *CircleQueue) AddQueue(val int) (err error) {
	if this.IsFull() {
		return errors.New("Queue is full!")
	}
	this.array[this.rear] = val //把值给尾部
	this.rear = (this.rear + 1) % this.maxsize
	return
}

func (this *CircleQueue) GetQueue() (val int, err error) {
	if this.IsEmpty() {
		err = errors.New("Queue is empty!")
	}
	val = this.array[this.front]
	this.front = (this.front + 1) % this.maxsize
	return
}

//判断环形队列满与空
func (this *CircleQueue) IsFull() bool {
	return (this.rear+1)%this.maxsize == this.front
}

func (this *CircleQueue) IsEmpty() bool {
	return this.rear == this.front
}

//取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	return (this.rear + this.maxsize - this.front) % this.maxsize
}

func (this *CircleQueue) ShowQueue() {
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	tempfront := this.front
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempfront, this.array[tempfront])
		tempfront = (tempfront + 1) % this.maxsize
	}
	fmt.Println()
}

func main() {
	//先创建一个队列
	circlequeue := &CircleQueue{
		maxsize: 5,
		rear:    0,
		front:   0,
	}

	//添加数组
	var key string
	var val int
	for {
		fmt.Println("1、输入add表示添加数据到队列")
		fmt.Println("2、输入get表示从队列中获取数据")
		fmt.Println("3、输入show表示显示队列")
		fmt.Println("4、输入exit表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入你要加入队列的数字：")
			fmt.Scanln(&val)
			err := circlequeue.AddQueue(val)
			if err != nil {
				fmt.Println("error：", err)
			} else {
				fmt.Println("添加成功！")
			}
		case "get":
			val, err := circlequeue.GetQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("从队列中取出：", val)
			}
		case "show":
			circlequeue.ShowQueue()
		case "exit":
			return
		default:
			fmt.Println("输入有误，请重新输入!!!")
		}

	}
}
