package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println(time.Now())
	timer2 := time.NewTimer(5*time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Timer 2 结束了。。。开始。。。。")
		fmt.Println(time.Now())
	}()

	time.Sleep(7*time.Second)
	/*flag := timer2.Stop()
	if flag{
		fmt.Println("Timer 2 停止了。。。")
	}*/
}
