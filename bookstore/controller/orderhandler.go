package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"net/http"
	"text/template"
	"time"
)

//Checkout 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户id
	UserID := session.UserID
	//获取购物车
	cart, _ := dao.GetCartByUserID(UserID)
	//生成订单好
	orderID := utils.CreateUUID()
	//创建Order
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(UserID),
	}
	//将订单保存到数据库
	dao.AddOrder(order)
	//保存订单项
	//获取购物车里的所有购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _, v := range cartItems {
		//创建订单项
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderID,
		}
		//保存购物项
		dao.AddOrderItem(orderItem)
		//更新当前购物项对应图书的库存和销量
		book := v.Book
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		//更新图书的信息
		dao.UpdateBook(book)
	}
	//清空购物车
	dao.DeleteCartByCartID(cart.CartID)
	//将订单号设置到session中
	session.Order = order
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, session)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	//调用dao中获取所有订单的函数
	orders, _ := dao.GetOrders()
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	//调用dao中获取所有订单的函数
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

func GetMyOrder(w http.ResponseWriter,r *http.Request){
	_,session := dao.IsLogin(r)
	//获取我的ID
	userID := session.UserID
	//获取订单
	orders,_ := dao.GetMyOrders(userID)
	//将订单设置到session中
	session.Orders = orders

	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)

}

//SendOrder 发货
func SendOrder(w http.ResponseWriter,r *http.Request){
	orderID := r.FormValue("orderId")
	//发货
	dao.UpdateOrderState(orderID,1)
	//查询所有的订单
	GetOrders(w,r)
}

//TakeOrder 发货
func TakeOrder(w http.ResponseWriter,r *http.Request){
	orderID := r.FormValue("orderId")
	//发货
	dao.UpdateOrderState(orderID,2)
	//查询我的订单
	GetMyOrder(w,r)
}
