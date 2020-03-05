package main

import (
	"fmt"
	//"time"
)

// func main() {
// 	ch := make(chan int)
// 	go hello(ch)
// 	//for{
// 	//	time.Sleep(1*time.Second)
// 	//	v, ok := <- ch
// 	//	if !ok{
// 	//		fmt.Println("已经读取了所有的数据。。",ok,v)
// 	//		break
// 	//	}
// 	//	fmt.Println("读取的数据：",v,ok)
// 	//}
// 	for v := range ch{
// 		time.Sleep(1*time.Second)
// 		fmt.Println("读取的数据：",v)
// 	}
// }

func hello(ch chan int){
	fmt.Println("发送数据流")
	for i := 1;i <=10;i++{
		ch <- i
	}
	close(ch)
	fmt.Println("通道关闭。。。")
}
