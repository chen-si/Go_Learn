package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go fun1()
	go fun2()

	fmt.Println("main等待。。。。。。")
	wg.Wait()
	fmt.Println("main...OVER...")
}

func fun1(){
	defer wg.Done()
	for i := 1;i <= 100; i++{
		fmt.Println("func1打印",i)
	}
}
func fun2(){
	defer wg.Done()
	for i := 1;i <= 100; i++{
		fmt.Println("\tfunc2打印",i)
	}
}
