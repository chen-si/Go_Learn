package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go fun1(ch1)
	go fun2(ch2)

	ch2 <- 200

	data := <- ch1


	fmt.Println("从fun1接收到",data)

	fmt.Println("main over....")
}

func fun1(ch chan int){
	fmt.Println("使用通道发送数据")
	//ch <- 100
}

func fun2(ch chan int){
	fmt.Println("使用通道接收数据")
	data := <- ch
	fmt.Println("从主goroutine中接收到",data)
}

