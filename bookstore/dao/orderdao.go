package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//AddOrder 向数据库中插入订单
func AddOrder(order *model.Order) error {
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	return err
}
