package model

//Cart 购物车
type Cart struct {
	CartID      string      //购物车ID
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //购物车里的图书总数量
	TotalAmount float64     //购物车的图书总金额
	UserID      int         //购物车所有者
	UserName    string
}

//GetTotalCount 获取图书总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//GetTotalAmount 获取图书总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
