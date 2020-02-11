package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string, 10)
	go sendData(ch)

	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("读完了。。", ok)
			break
		}
		fmt.Println("\t读取的数据是：", v)
	}
	fmt.Println("main。。over。。。。")

}

func sendData(ch chan  string){
	for i:= 0;i<10;i++{
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine中写出第 %d 个数据\n",i)
	}
	close(ch)
}
