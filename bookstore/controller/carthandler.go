package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"net/http"
)

//AddBook2Cart 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		//已经登录
		bookId := r.FormValue("bookId")
		//根据图书id获取图书信息
		book, err := dao.GetBookByID(bookId)
		if err != nil {
			fmt.Println(err)
		}
		//判断是否登录

		//获取用户的id
		UserID := session.UserID
		//判断购物车中是否有当前用户的购物车
		cart, _ := dao.GetCartByUserID(UserID)
		if cart != nil {
			//已经有购物车了  此时需要判断购物车是否有当前图书
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookId, cart.CartID)
			if cartItem != nil {
				//已经有这一购物项了 购物项的数量加一即可
				//获取购物车切片中的所有购物项
				cartItems := cart.CartItems
				//遍历得到每一个购物项
				for _, v := range cartItems {
					//寻找匹配当前书籍的购物项
					if v.Book.ID == cartItem.Book.ID {
						//将当前购物项中的图书数量加一
						v.Count++
						//将数据库中该购物项更新
						dao.UpdateBookCount(v.Count, v.Book.ID, cart.CartID)
					}
				}
			} else {
				//创建购物项
				cartItem = &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				//将购物项添加到倩倩购物车的切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将新创建的购物项添加到数据库中
				dao.AddCartItem(cartItem)
			}
			//不管之前购物车中是否有当前图书对应的购物项，都需要更新购物车中图书的总数量和总金额
			dao.UpdateCart(cart)

		} else {
			//当前用户还没有购物车，需要创建购物车并添加到数据库
			//创建购物车
			//生成购物车的id
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: UserID,
			}
			//创建购物项
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}
			cartItems = append(cartItems, cartItem)

			cart.CartItems = cartItems
			//将购物车保存到数据库中
			dao.AddCart(cart)
		}
		//显示将谁加入到了购物车中
		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车"))
	} else {
		//没有登录
		w.Write([]byte("请先登录！"))
	}

}
