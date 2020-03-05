// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// var ticket = 10
// var wait sync.WaitGroup
// var mt sync.Mutex

// func main() {

// 	wait.Add(4)
// 	go saleTickets("售票口1")
// 	go saleTickets("售票口2")
// 	go saleTickets("售票口3")
// 	go saleTickets("售票口4")

// 	wait.Wait()
// }

// func saleTickets(name string) {
// 	defer wait.Done()
// 	rand.Seed(time.Now().UnixNano())
// 	for {
// 		mt.Lock()
// 		if ticket > 0 {
// 			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// 			fmt.Println(name, "售出：", ticket)
// 			ticket--
// 		} else {
// 			mt.Unlock()
// 			fmt.Println(name, "售罄，没有票了。。")
// 			break
// 		}
// 		mt.Unlock()
// 	}
// }
