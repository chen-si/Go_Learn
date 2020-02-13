package main

import (
	"errors"
	"fmt"
)

//使用一个结构体管理队列
type Queue struct {
	maxsize int
	array   [5]int
	front   int //队首
	rear    int //队尾
}

func (this *Queue) AddQueue(val int) error {
	//1、先判断队列是否已满
	if this.rear == this.maxsize-1 {
		return errors.New("Queue full")
	}

	this.rear++ //rear 后移
	this.array[this.rear] = val

	return nil
}

func (this *Queue) GetQueue() (val int, err error) {
	if this.rear == this.front {
		err = errors.New("Queue is empty")
		return -1, err
	}
	this.front++
	val = this.array[this.front]
	return
}

func (this *Queue) ShowQueue() {
	fmt.Println("当前队列是：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("%d\t", this.array[i])
	}
	fmt.Println()
}

func main() {
	//先创建一个队列
	queue := &Queue{
		maxsize: 5,
		rear:    -1,
		front:   -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("error：", err)
			} else {
				fmt.Println("添加成功！")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("从队列中取出：", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			return
		default:
			fmt.Println("输入有误，请重新输入!!!")
		}

	}
}
