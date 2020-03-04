package model

//Order 结构
type Order struct {
	OrderID     string  //订单号
	CreateTime  string  //生成订单的时间
	TotalCount  int64   //订单中图书的总数量
	TotalAmount float64 //总金额
	State       int64   //0 未发货 1 已发货 2交易完成
	UserID      int64   //订单的所属用户
}

//是否发货
func (order *Order) NoSend() bool {
	return order.State == 0
}

//已发货
func (order *Order) SendComplete() bool {
	return order.State == 1
}

//交易完成
func (order *Order) Complete() bool {
	return order.State == 2
}
