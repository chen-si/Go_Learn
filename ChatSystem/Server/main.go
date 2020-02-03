package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){

}

func main() {
	fmt.Println("服务器开始监听8889端口......")

	listener,err :=net.Listen("tcp","0.0.0.0:8889")
	if err != nil{
		fmt.Println("Linten errpr:",err)
		return
	}
	defer listener.Close()
	for{
		fmt.Println("等待客户端连接服务器......")
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println("listener Accept error:",err)
		}
		//goroutine处理不同服务器的请求
		go process(conn)
	}
}