package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	//设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//主页
	http.HandleFunc("/main", controller.IndexHandler)
	//登录
	http.HandleFunc("/login", controller.Login)
	//注册
	http.HandleFunc("/regist", controller.Regist)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	//http.HandleFunc("/getBooks", controller.GetBooks)
	//获取带分页的图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//按价格范围获取图书
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	// //添加图书
	// http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//去更新图书的页面
	http.HandleFunc("/toupdateBookPage", controller.ToUpdateBookPage)
	//更新或添加图书
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)

	http.ListenAndServe(":8080", nil)
}
