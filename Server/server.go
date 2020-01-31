package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	fmt.Printf("Client: %s\n",conn.RemoteAddr().String())
	for{
		//循环读取来自Client的信息
		buf := make([]byte,1024)
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("Read error:",err)
			break
		} else{
			s := string(buf[:n])
			fmt.Println("Receive from Client：",s)

			//向客户端返回信息
			if s == "hello"{
				_,err = conn.Write([]byte("hello,liu"))
				if err != nil{
					fmt.Println("Server write error:",err)
				}
			}else{
				_,err = conn.Write([]byte("Sorry,I can't understand"))
				if err != nil{
					fmt.Println("Server write error:",err)
				}
			}
		}

	}
}

func main() {
	fmt.Println("Server start....")
	listenner,err := net.Listen("tcp","127.0.0.1:8888")
	if err != nil{
		fmt.Println("Server start error:",err)
		return
	}
	defer listenner.Close()
	for{
		fmt.Println("Server are waiting for connect...")
		conn,err := listenner.Accept()
		if err != nil{
			fmt.Println("conn error:",err)
		}else{
			go process(conn)
		}

	}
}
