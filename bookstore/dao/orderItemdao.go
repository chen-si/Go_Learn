package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//AddOrderItem 向数据库中插入订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	sqlStr := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	return err
}

//GetOrderItemsByOrderID 根据订单号获取所有订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	sqlStr := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id = ?"
	rows, err := utils.Db.Query(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		//添加到切片中
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
