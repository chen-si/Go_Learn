package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//AddCartItem 向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"

	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByBookIDAndCartID 根据图书的ID 获取对应的购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	//得到图书
	book, _ := GetBookByID(bookID)
	//创建一个CartItem
	cartItem := &model.CartItem{}
	cartItem.Book = book
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

//GetCartItemsByCartID 根据购物车ID 获取购物车对应的所有购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		//设置一个变量接受bookid
		var bookID string
		cartItem := &model.CartItem{}
		err = rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err != nil {
			return nil, err
		}
		book, _ := GetBookByID(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//UpdateBookCount 根据图书id和购物车id以及图书的数量更新购物项
func UpdateBookCount(cartItem *model.CartItem) error {
	sqlStr := "update cart_items set count = ?,amount = ? where book_id = ? and cart_id = ?"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	return err
}

//DeleteCartItemsByCartID 根据购物车ID删除对应的所有购物项
func DeleteCartItemsByCartID(cartID string)error{
	sqlStr := "delete from cart_items where cart_id = ?"
	_,err := utils.Db.Exec(sqlStr,cartID)
	return err
}

//DeleteCartItemByID 根据购物项的ID删除购物项
func DeleteCartItemByID(cartItemID string) error{
	sqlStr := "delete from cart_items where id = ?"
	_,err := utils.Db.Exec(sqlStr,cartItemID)
	return err
}
