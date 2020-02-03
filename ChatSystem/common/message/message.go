package common

//消息类型常量表示
const (
	LoginMesType		="LoginMes"
	LoginResMesType		="LoginResMes"
)


//Message 的一般类型
type Message struct{
	Type 	string//消息类型
	Data 	string//消息内容
}

//登录消息
type LoginMes struct{
	UserID 		int		//用户ID
	UserPWD 	string	//用户密码
	UserName 	string	//用户名
}

//登录结果消息
//code：
//500 ：未注册
//200 ：登陆成功
type LoginResMes struct{
	code int//返回状态码
	Error string//错误消息
}


