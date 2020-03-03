package model

//OrderItem 订单项结构
type OrderItem struct {
	OrderItemID int64   //订单项的ID
	Count       int64   //数量
	Amount      float64 //金额
	Title       string  //图书名
	Author      string  //作者
	Price       float64 //价格
	ImgPath     string  //图书封面
	OrderID     string  //订单项所属的订单
}
