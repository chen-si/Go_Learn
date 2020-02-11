package main

import (
	process2 "ChatSystem/Server/process"
	"ChatSystem/Server/utils"
	"ChatSystem/common/message"
	"fmt"
	"io"
	"net"
)

//先创建Processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个serverProcessMes函数处理不同的消息
//根据消息类型的而不同调用不同的函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//看看能否从客户端接收到smsMes
	//fmt.Println(mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理用户登录
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)

	default:
		fmt.Println("不支持的消息类型...")

	}
	return
}

func (this *Processor) Process3() (err error) {
	//接收消息
	for {
		//将读取数据包封装成一个函数
		//创建一个Transfer实例+++++++++++++++++++++++++++++++++++++++++
		tf := utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出...")
				return err
			}
			fmt.Println("readPkg(conn) error:", err)
			return err
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println(" serverProcessMes(conn,&mes) error:", err)
			return err
		}
	}
}
