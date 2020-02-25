package controllers

import (
	"fmt"
	"web/models"

	"github.com/astaxie/beego"
)

//由model这个名字叫UserInfo 操作的表确实是user_info
type UserInfo struct {
	Id       int64
	Username string
	Password string
}

type TestModelController struct {
	beego.Controller
}

func (c *TestModelController) Get() {
	// //链接数据库
	// orm.RegisterDataBase("default", "mysql", "root:1234w5asd@tcp(localhost:3306)/test?charset=utf8", 30)
	// //绑定表
	// orm.RegisterModel(new(UserInfo))

	// o := orm.NewOrm()

	//插入
	// user := UserInfo{
	// 	Username: "张三",
	// 	Password: "123456",
	// }
	// id, err := o.Insert(&user)
	//更新
	// user := UserInfo{
	// 	Username: "张三",
	// 	Password: "123456",
	// }
	// user.Id=1
	// user.Username="李四"
	// id,err := o.Update(&user)
	//read
	// user := UserInfo{
	// 	Id: 1,
	// }
	// o.Read(&user)
	//下面是原生读取
	// var maps []orm.Params
	// o.Raw("select * from user_info").Values(&maps)

	// for _, v := range maps {
	// 	c.Ctx.WriteString(fmt.Sprintf("%v", v))
	// }

	// //采用QueryBuilder的方式进行读取
	// var users []UserInfo
	// qb, _ := orm.NewQueryBuilder("mysql")

	// qb.Select(" * ").From("user_info").Where("username= ?")

	// sql := qb.String()
	// o.Raw(sql, "李四").QueryRows(&users)
	// c.Ctx.WriteString(fmt.Sprintf("%v", users))

	user := &models.UserInfo{
		Username: "wangwu",
		Password: "123456789",
	}
	id, err := models.AddUser(user)
	c.Ctx.WriteString(fmt.Sprintf("id:%v, err:%v", id, err))
}
