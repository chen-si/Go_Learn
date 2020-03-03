package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
	"time"
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

// func TestBook(t *testing.T) {
// 	fmt.Println("测试bookdao中的相关函数")
// 	//t.Run("测试获取所有图书", testGetBooks)
// 	// t.Run("测试添加图书", testAddBooks)
// 	// t.Run("测试删除图书", testDeleteBooks)
// 	//t.Run("测试获取图书", testGetBookByID)
// 	//t.Run("测试更新图书", testUpdateBook)
// 	// t.Run("测试带分页的图书", testGetPageBooks)
// 	t.Run("测试带分页和价格范围的图书", testGetPageBooksByPrice)
// }
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

func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("9")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.Totalrecord)
	fmt.Println("当前页的图书有是：")
	for k, v := range page.Books {
		fmt.Printf("第%d本图书是%v\n", k+1, v)
	}
}
func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("3", "10", "30")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.Totalrecord)
	fmt.Println("当前页的图书有是：")
	for k, v := range page.Books {
		fmt.Printf("第%d本图书是%v\n", k+1, v)
	}
}

// func TestSession(t *testing.T) {
// 	fmt.Println("开始测试Session相关函数")
// 	//t.Run("测试添加", testAddSession)
// 	//t.Run("测试删除", testDeleteSession)
// 	t.Run("测试查询", testGetSession)
// }

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "13838381438",
		UserName:  "张三",
		UserID:    5,
	}
	AddSession(sess)
}
func testDeleteSession(t *testing.T) {
	DeleteSession("13838381438")
}

func testGetSession(t *testing.T) {
	sess, _ := GetSession("5972e1d6-2eb6-4e7a-73cf-5e015e98d21b")
	fmt.Println(sess)
}

// func TestCart(t *testing.T) {
// 	fmt.Println("测试购物车相关函数")
// 	// t.Run("测试添加购物车", testAddCart)
// 	// t.Run("测试根据图书id获取购物项", testGetCartItemByBookID)
// 	// t.Run("测试根据购物车id获取所有购物项", testGetCartItemsByCartID)
// 	//t.Run("测试根据用户id获取购物车", testGetCartByUserID)
// 	//t.Run("测试根据图书id和购物车id以及图书的数量更新购物项", testUpdateBookCount)
// 	//t.Run("测试根据购物车id删除购物车", testDeleteCartByCartID)
// 	t.Run("测试根据购物项id删除购物项", testDeleteCartItemByID)
// }

func testAddCart(t *testing.T) {
	//设置要买的第一本数
	book := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	//创建两个购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartID: "66668888",
	}
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "66668888",
	}
	//创建购物车切片
	var catrItems []*model.CartItem
	catrItems = append(catrItems, cartItem)
	catrItems = append(catrItems, cartItem2)

	//创建购物车
	cart := &model.Cart{
		CartID:    "66668888",
		CartItems: catrItems,
		UserID:    1,
	}
	//测试将购物车插入数据库中
	err := AddCart(cart)
	fmt.Println(err)
}

func testGetCartItemByBookID(t *testing.T) {
	cartItem, err := GetCartItemByBookIDAndCartID("1", "66668888")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cartItem)
}

func testGetCartItemsByCartID(t *testing.T) {
	cartItems, err := GetCartItemsByCartID("66668888")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range cartItems {
		fmt.Println(v)
	}
}

func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(1)
	fmt.Println(cart)
	for _, v := range cart.CartItems {
		fmt.Println(v)
	}
}

// func testUpdateBookCount(t *testing.T) {
// 	err := UpdateBookCount(15, 1, "66668888")
// 	fmt.Println(err)
// }

func testDeleteCartByCartID(t *testing.T) {
	DeleteCartByCartID("58b291f0-db22-4073-4930-816959072bb9")
}

func testDeleteCartItemByID(t *testing.T) {
	err := DeleteCartItemByID("3")
	fmt.Println(err)
}

func TestOrder(t *testing.T) {
	fmt.Println("测试订单相关函数")
	t.Run("测试添加订单和订单项", testAddOrder)
}
func testAddOrder(t *testing.T) {
	//创建订单
	order := &model.Order{
		OrderID:     "13811118888",
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      1,
	}
	//创建订单项
	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  300,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   300,
		ImgPath: "/static/img/default.jpg",
		OrderID: "13811118888",
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderID: "13811118888",
	}
	//保存订单
	AddOrder(order)
	//保存订单项
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
}
