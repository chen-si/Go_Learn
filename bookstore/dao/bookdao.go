package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//GetBooks 获取数据库中所有图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		//var book *model.Book
		book := &model.Book{}
		//个体book字段复制
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

//AddBook 向数据库中添加图书
func AddBook(b *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook 根据图书id删除图书
func DeleteBook(bookID string) error {
	sqlStr := "delete from books where id = ?"
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

//GetBookByID 根据图书的id从数据库查询出一本图书
func GetBookByID(bookID string) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	row := utils.Db.QueryRow(sqlStr, bookID)

	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

//UpdateBook 根据图书的ID更新图书信息
func UpdateBook(b *model.Book) error {
	//sql
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	//执行
	_ ,err := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ID)
	if err != nil{
		return err
	}
	return nil
}
