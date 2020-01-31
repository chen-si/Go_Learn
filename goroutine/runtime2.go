package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	//获取逻辑cpu的数量
	fmt.Println("逻辑CPU的数量-->",runtime.NumCPU())

	//设置go程序执行的最大的cpu的数量：[1,256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)
}
func main() {
	go func() {
		fmt.Println("goroutine开始。。")
		//调用fun
		fun()
		fmt.Println("goroutine结束。。。")
	}()
	time.Sleep(3*time.Second)
}

func fun(){
	defer fmt.Println("derfer...")
	//return
	runtime.Goexit()
	fmt.Println("fun函数。。。")
}