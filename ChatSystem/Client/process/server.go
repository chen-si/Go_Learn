package process

import (
	"ChatSystem/Client/utils"
	"ChatSystem/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面

func ShowMenu() {
	fmt.Println("---------恭喜xxx登陆成功---------")
	fmt.Println("---------1、显示在线用户列表-----")
	fmt.Println("---------2、发送消息-------------")
	fmt.Println("---------3、消息列表-------------")
	fmt.Println("---------4、退出系统-------------")
	fmt.Println("请选择（1-4）：")
	var key int
	var content string
	//因为我们总会使用到SmsProcess 因此将其定义在switch外部
	smsProcess := &SmsProcess{}

	_, _ = fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUsers()
	case 2:
		fmt.Println("请输入你想对大家说的话")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("消息列表")
		
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("")

	}

}

//和服务器段保持通讯
func ServerProcessMes(Conn net.Conn) {
	//创建一个transfer实例 循环读取消息
	tf := &utils.Transfer{
		Conn: Conn,
	}
	for {
		fmt.Printf("客户端：%s正在等待读取服务器发送的消息...\n", Conn.LocalAddr().String())
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("ServerProcessMes tf.ReadPkg() error:", err)
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了
			//1、取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2、把这个人加入到客户端维护的map[int]User中去
			UpdateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回了一个未知的消息类型")

		}

	}
}
