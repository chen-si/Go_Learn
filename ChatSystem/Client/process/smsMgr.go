package process

import (
	"ChatSystem/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	//显示即可
	//1、反序列化mes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes error:", err)
		return
	}

	//显示信息
	info := fmt.Sprintf("用户%d对大家说：%s", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
}
