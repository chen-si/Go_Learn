package process2

import (
	"ChatSystem/Server/model"
	"ChatSystem/Server/utils"
	"ChatSystem/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段 表示是哪一个用户的conn
	UserID int
}

//这里我们编写通知用户上线的方法
//userid 通知其他在线用户 我上线了
func (this *UserProcess) NotifyOtherOnlineUser(UserID int) {
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == UserID {
			continue
		}
		//开始通知【单独的一个函数】
		up.NotifyMeOnline(UserID)
	}
}

func (this *UserProcess) NotifyMeOnline(UserID int) {
	//组装我们的消息NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = UserID
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) error:", err)
		return
	}

	//把序列化的notifyUserStatusMes赋值给 mes.Data
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) error:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline tf.WritePkg(data) error:", err)
		return
	}

}

//编写一个serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//处理登录
	//先从mes.Data反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),loginMes) error:", err)
		return err
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	//到redis数据库验证用户
	//使用model.MyUserDao 到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPWD)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200

		//将登录成功的用户id赋给this
		this.UserID = loginMes.UserID
		//把登录成功的用户放入userMgr中
		userMgr.AddOnlineUser(this)
		//通知其他在线的用户 我上线了
		this.NotifyOtherOnlineUser(loginMes.UserID)

		//遍历userid
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID, id)
		}
		fmt.Println(user, "登陆成功...")
	}

	// //如果用户名为100 密码为123456表示登录成功
	// if loginMes.UserID == 100 && loginMes.UserPWD == "123456"{
	// 	//登陆成功
	// 	loginResMes.Code=200 //200表示登录成功

	// }else{
	// 	//登录失败
	// 	loginResMes.Code=500 //500表示用户未注册
	// 	loginResMes.Error="该用户不存在，请注册再使用"
	// }
	//反序列化loginResMes并复制给resMes.Data
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) error", err)
		return err
	}
	resMes.Data = string(data)

	//反序列化resMes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) error", err)
		return err
	}
	//发送data 将发送函数封装成writePkg（）
	//因为使用了分层模式（mvs）我们先创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return err
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),&registerMes) error:", err)
		return err
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//到redis数据库注册用户
	//使用model.MyUserDao 到redis去注册
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功...")
	}

	//反序列化registerResMes并复制给resMes.Data
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) error", err)
		return err
	}
	resMes.Data = string(data)

	//反序列化resMes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) error", err)
		return err
	}
	//发送data 将发送函数封装成writePkg（）
	//因为使用了分层模式（mvs）我们先创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}
