package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go SendData(ch1)
	for v := range ch1{
		time.Sleep(1*time.Second)
		fmt.Println("接收到的数据为",v)
	}

	ch2 := make(chan int)
	go RecieveData(ch2)
	ch2 <- 99

	fmt.Println("main...over...")

}

func SendData(ch chan<- int){
	fmt.Println("发送数据的单向通道")
	for i := 1;i<= 10 ;i++{
		ch <- i
	}
	close(ch)
}

func RecieveData(ch <-chan int){
	fmt.Println("接收数据的单向通道")
 	data := <-ch
 	fmt.Println("接收到的数据为",data)
}
