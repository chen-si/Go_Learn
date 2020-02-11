package main

import (
	"ChatSystem/Client/utils"
	"ChatSystem/common/message"
	"encoding/json"
	"fmt"
	"net"
)

func login(UserID int,UserPWD string) (err error){
	//1、链接服务器
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error:",err)
		return
	}
	defer conn.Close()

	//2、准备通过conn发消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3、创建login message结构体
	var loginMes message.LoginMes
	loginMes.UserID = UserID
	loginMes.UserPWD = UserPWD

	//4、将loginMes 序列化
	data,err :=json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal(loginMes) error:",err)
		return
	}

	//5、把data复制给mes.Data
	mes.Data = string(data)

	//6、把 mes 序列化
	data,err = json.Marshal(mes)
	if err!= nil{
		fmt.Println(" json.Marshal(mes) error:",err)
		return
	}
	//7、发送数据data
	//7.1 、先把data的长度发送给服务器
	//先获取到data的长度，然后转化为一个表示长度的切片
	err = utils.writePkg(conn,data)
	if err != nil{
		fmt.Println(" writePkg(conn,data) error",err)
		return
	}
	//处理服务器段返回的消息
	mes,err = utils.readPkg(conn)
	if err != nil{
		fmt.Println("readPkg(conn) error:",err)
		return
	}
	//将返回的mes.Data反序列化成loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if err != nil{
		fmt.Println("json.Unmarshal([]byte(mes.Data),loginResMes) error:",err)
		return
	}
	if loginResMes.Code == 200{
		fmt.Println("登录成功")
	}else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}


	return
}
