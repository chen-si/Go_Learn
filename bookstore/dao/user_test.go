package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试bookdao中的方法")
	m.Run()
}

// func TestUser(t *testing.T) {
// 	// fmt.Println("开始测试userdao中的函数")
// 	// t.Run("验证用户名或密码", testLogin)
// 	// t.Run("验证用户名", testRegister)
// 	// t.Run("验证保存用户", testSaveUser)
// }

// func testLogin(t *testing.T) {
// 	user, _ := CheckUsernameAndPassword("admin", "123456")
// 	fmt.Println(user)
// }

// func testRegister(t *testing.T) {
// 	user, _ := CheckUsername("admin")
// 	fmt.Println(user)
// }

// func testSaveUser(t *testing.T) {
// 	SaveUser("admin3", "123456", "admin@liu.com")
// }

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有图书", testGetBooks)
	// t.Run("测试添加图书", testAddBooks)
	// t.Run("测试删除图书", testDeleteBooks)
	//t.Run("测试获取图书", testGetBookByID)
	t.Run("测试更新图书", testUpdateBook)
}
func testGetBooks(t *testing.T) {
	books, _ := GetBooks()

	for k, v := range books {
		fmt.Printf("第%d本图书的信息是：%v\n", k+1, v)
	}
}

func testAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "红楼梦",
		Author:  "曹雪芹",
		Price:   77.77,
		Sales:   100,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}
	//调用添加图书的函数
	_ = AddBook(book)
}

func testDeleteBooks(t *testing.T) {
	DeleteBook("33")
}

func testGetBookByID(t *testing.T) {
	book, _ := GetBookByID("21")
	fmt.Println(book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      32,
		Title:   "3个女人与105个男人的故事",
		Author:  "施耐庵",
		Price:   77.77,
		Sales:   10000,
		Stock:   1,
		ImgPath: "static/img/default.jpg",
	}
	UpdateBook(book)
}
