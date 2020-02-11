package main

import "fmt"

var UserID int
var UserPWD string

func main() {
	fmt.Println("----------------------欢迎登录聊天系统----------------------")
	fmt.Printf("\t\t\t 1 登录系统\n")
	fmt.Printf("\t\t\t 2 注册账号\n")
	fmt.Printf("\t\t\t 3 退出\n")
	fmt.Printf("\t\t\t 请选择（1-3）\n")
	//接收用户的输入的选择
	var key int
	//判断循环是否继续进行
	var loop = true

	for loop {
		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("开始登录")
			loop = false
		case 2:
			fmt.Println("开始注册")
			loop = false
		case 3:
			fmt.Println("退出成功...")
			loop = false
		default:
			fmt.Println("输入有误，请重新选择功能")
		}
	}
	fmt.Println("---------------------------------------------------------")

	if key ==1 {
		fmt.Println("请输入账户名：")
		_, _ = fmt.Scanf("%d\n",&UserID)
		fmt.Println("请输入密码：")
		_, _ = fmt.Scanf("%s\n",&UserPWD)

		err := login(UserID,UserPWD)
		fmt.Println(err)
	}
}
