package model

//CartItem 购物项
type CartItem struct {
	CartItemID int64   //购物项ID
	Book       *Book   //购物项的图书信息
	Count      int64   //购物项的图书数量
	Amount     float64 //购物项的图书金额小计
	CartID     string  //购物车ID 通过uuid生成
}

//GetAmount  获取购物项中的金额小计
func (cartItem *CartItem)GetAmount()float64{
	//获取当前图书的价格
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
