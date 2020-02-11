package process

import (
	"ChatSystem/Client/utils"
	"ChatSystem/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
	//暂时不需要字段
}

func (this *UserProcess) Register(UserID int, UserPWD string,
	UserName string) (err error) {
	//1、链接服务器
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	//3、创建login message结构体
	var registerMes message.RegisterMes
	registerMes.User.UserID = UserID
	registerMes.User.UserPWD = UserPWD
	registerMes.User.UserName = UserName

	//4、将registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) error:", err)
		return
	}

	//5、把data复制给mes.Data
	mes.Data = string(data)

	//6、把 mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(" json.Marshal(mes) error:", err)
		return
	}

	//7、发送数据data
	//7.1 、先把data的长度发送给服务器
	//先获取到data的长度，然后转化为一个表示长度的切片
	//先创建一个Transfer实例

	tf := &utils.Transfer{
		Conn: conn,
	}
	fmt.Println(string(data))

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(" writePkg(conn,data) error", err)
		return
	}

	//处理服务器段返回的消息
	mes, err = tf.ReadPkg() //mes.Type = RegisterResMesType
	if err != nil {
		fmt.Println("readPkg(conn) error:", err)
		return
	}

	//将返回的mes.Data反序列化成registerResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),registerResMes) error:", err)
		return
	}

	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return
}

func (this *UserProcess) Login(UserID int, UserPWD string) (err error) {
	//1、链接服务器
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial error:", err)
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
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) error:", err)
		return
	}

	//5、把data复制给mes.Data
	mes.Data = string(data)

	//6、把 mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(" json.Marshal(mes) error:", err)
		return
	}
	//7、发送数据data
	//7.1 、先把data的长度发送给服务器
	//先获取到data的长度，然后转化为一个表示长度的切片
	//先创建一个Transfer实例

	tf := &utils.Transfer{
		Conn: conn,
	}
	fmt.Println(string(data))

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(" writePkg(conn,data) error", err)
		return
	}
	//处理服务器端返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) error:", err)
		return
	}
	//将返回的mes.Data反序列化成loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),loginResMes) error:", err)
		return
	}
	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")

		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserID = UserID
		CurUser.UserStatus = message.UserOnline


		//现在可以显示当前在线用户的列表
		fmt.Println("当前用户在线列表如下：")
		for _, id := range loginResMes.UsersID {
			if id == loginMes.UserID {
				continue
			}
			fmt.Println("用户ID = ", id)

			user := &message.User{
				UserID:     id,
				UserStatus: message.UserOnline,
			}

			onlineUsers[id] = user
		}
		fmt.Printf("\n\n")
		//显示菜单 循环显示
		//这里我们起一个协程
		//与服务器保持通讯，如果服务器有消息推送给客户端
		//则显示在客户端的终端

		go ServerProcessMes(conn)

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
