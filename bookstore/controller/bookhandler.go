package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

//去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中得到带分页图书的函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w, page)
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中得到带分页图书的函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, page)

}

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用bookdao中得到带分页图书的函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//调用bookdao中得到带分页图书的函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	//调用IsLogin函数判断是否已经登录
	flag, session := dao.IsLogin(r)
	if flag {
		//已经登录设置page和Username中的字段
		page.IsLogin = true
		page.Username = session.UserName
	}

	//解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}

// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	//调用bookdao中得到所有图书的函数
// 	books, _ := dao.GetBooks()
// 	//解析模板文件
// 	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
// 	//执行
// 	t.Execute(w, books)

// }

// //AddBook 添加图书
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	//获取用户输入的图书信息
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")
// 	//将价格 销量 和库存进行转换
// 	fPrice, _ := strconv.ParseFloat(price, 64)
// 	iSales, _ := strconv.ParseInt(sales, 10, 0)
// 	iStock, _ := strconv.ParseInt(stock, 10, 0)

// 	book := &model.Book{
// 		Title:   title,
// 		Author:  author,
// 		Price:   fPrice,
// 		Sales:   int(iSales),
// 		Stock:   int(iStock),
// 		ImgPath: "static/img/default.jpg",
// 	}
// 	//调用添加图书的函数
// 	dao.AddBook(book)
// 	//调用GetBooks函数 再查询一遍数据库
// 	GetBooks(w, r)
// }

//DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookid")
	dao.DeleteBook(bookID)
	//再次查询一遍数据库
	GetPageBooks(w, r)
}

//ToUpdateBookPage 去更新或者添加图书图书
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookid")
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 {
		//在更新图书
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		//在添加 图书
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
	//解析模板

}

//UpdateBook 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的图书信息
	bookID := r.PostFormValue("bookid")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//将价格 销量 和库存进行转换
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	ibookID, _ := strconv.ParseInt(bookID, 10, 0)

	book := &model.Book{
		ID:      int(ibookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "static/img/default.jpg",
	}
	if book.ID > 0 {
		//更新
		dao.UpdateBook(book)
	} else {
		//添加
		dao.AddBook(book)
	}
	//调用GetBooks函数 再查询一遍数据库
	GetPageBooks(w, r)
}
