package model

import (
	"fmt"
	"testing"
)

//TestMain可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始:")
	//通过m.Run 执行测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	//通过t.Run执行子测试函数
	//t.Run("测试添加用户", testAddUser)
	//t.Run("测试获取用户", testGetUserById)
	t.Run("测试获取所有用户", tetsGetUsers)
}

//如果函数不以Test开头 默认不执行
//我们可以将他设置为一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户子测试函数执行")
	// user := &User{}

	//调用添加用户的方法
	// _ = user.AddUser()
	// _ = user.AddUser2()
}

func testGetUserById(t *testing.T) {
	fmt.Println("子测试，查询一条记录")
	user := User{
		ID: 1,
	}
	//调用方法
	U, _ := user.GetUserById()
	fmt.Println(U)

}

func tetsGetUsers(t *testing.T) {
	fmt.Println("子测试,查询所有记录")
	user := &User{}
	us, _ := user.GetUsers()
	for k, v := range us {
		fmt.Printf("第%d个用户是：%v\n", k+1, v)
	}
}
