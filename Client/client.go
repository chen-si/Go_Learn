package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {//127.0.0.1
	conn,err := net.Dial("tcp","localhost:8888")
	if err != nil{
		fmt.Println("conn error:",err)
		return
	}

	for{
		StringReader := bufio.NewReader(os.Stdin)
		s,err := StringReader.ReadString('\n')
		if err != nil{
			fmt.Println("Read string error:",err)
		}else {
			s = strings.Trim(s," \r\n")
			if s == "exit" {
				break
			}
			_,err := conn.Write([]byte(s))
			if err != nil{
				fmt.Println("Write Data error:",err)
			}else{
				buf := make([]byte,1024)
				n,err := conn.Read(buf)
				if err != nil{
					fmt.Println("Client Read error:",err)
					return
				}
				fmt.Println("Receive from Server:",string(buf[:n]))
			}
		}
	}
	fmt.Println("Client exit...")
}
