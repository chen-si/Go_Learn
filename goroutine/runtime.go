// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Println(runtime.GOROOT())

// 	fmt.Println(runtime.GOOS)

// 	fmt.Println(runtime.NumCPU())

// 	runtime.GOMAXPROCS(runtime.NumCPU())

// 	go printNum()

// 	for i:=0;i<4;i++{
// 		runtime.Gosched()
// 		fmt.Println("main...")
// 	}
// 	fmt.Println("main...over...")
// }

// func printNum() {

// 	for i := 1; i <= 20; i++ {
// 		fmt.Printf("子goroutine中打印数字：%d\n", i)
// 	}
// }