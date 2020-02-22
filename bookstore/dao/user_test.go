package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("开始测试userdao中的函数")
	t.Run("验证用户名或密码", testLogin)
	t.Run("验证用户名", testRegister)
	t.Run("验证保存用户", testSaveUser)
}

func testLogin(t *testing.T) {
	user, _ := CheckUsernameAndPassword("admin", "123456")
	fmt.Println(user)
}

func testRegister(t *testing.T) {
	user, _ := CheckUsername("admin")
	fmt.Println(user)
}

func testSaveUser(t *testing.T) {
	SaveUser("admin2", "123456", "admin@liu.com")
}
