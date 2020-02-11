package message

//消息类型常量表示
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//这里我们定义几个用户在线状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

//Message 的一般类型
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"date"` //消息内容
}

//登录消息
type LoginMes struct {
	UserID   int    `json:"userid"`   //用户ID
	UserPWD  string `json:"userpwd"`  //用户密码
	UserName string `json:"username"` //用户名
}

type LoginResMes struct {
	//登录结果消息
	//code：
	//500 ：未注册
	//200 ：登陆成功
	//403 ：密码错误
	//505 : 服务器内部错误
	Code    int    `json:"code"`    //返回状态码
	UsersID []int  `json:"usersid"` //保存用户id的切片
	Error   string `json:"error"`   //错误消息
}

type RegisterMes struct {
	//注册消息
	User User `json:"user"`
}

type RegisterResMes struct {
	//400表示用户已占用
	//200表示注册成功
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//为了配合服务器推送用户状态变化消息的类型
type NotifyUserStatusMes struct {
	UserID int `json:"userid"` //用户id
	Status int `json:"status"` //用户状态
}

//增加一个smsMes 发送消息
type SmsMes struct {
	Content string `json:"content"`
	User           //继承
}
