package process

import (
	"ChatSystem/Client/utils"
	"ChatSystem/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1、创建一个Message
	var mes message.Message
	mes.Type = message.SmsMesType

	//2、创建一个smsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content //内容
	smsMes.UserID = CurUser.UserID
	smsMes.UserStatus = CurUser.UserStatus

	//3、 序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(smsMes) error:", err)
		return
	}
	mes.Data = string(data)

	//4、把 mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(mes) error:", err)
		return
	}

	//5、发送数据data
	//5.1 、先把data的长度发送给服务器
	//先获取到data的长度，然后转化为一个表示长度的切片
	//先创建一个Transfer实例

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	fmt.Println(string(data))

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes writePkg(data) error", err)
		return
	}

	return
}
