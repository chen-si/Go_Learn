package process2

import (
	"ChatSystem/Client/utils"
	"ChatSystem/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	//暂时不需要字段
}

//转发消息的方法
func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	//遍历服务器端的onlineUsers map[int]*userprocess
	//将消息转发出去
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal([]byte(mes.Data),&smsMes) error:", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(mes) error:", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//这里我们需要过滤掉自己
		if id == smsMes.UserID {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
	return
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个transfer实例 发送data
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser error:", err)
		return
	}
}
