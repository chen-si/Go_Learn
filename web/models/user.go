package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
var db orm.Ormer
//由model这个名字叫UserInfo 操作的表确实是user_info
type UserInfo struct {
	Id       int64
	Username string
	Password string
}

func init() {
	orm.Debug = true
	//链接数据库
	orm.RegisterDataBase("default", "mysql", "root:1234w5asd@tcp(localhost:3306)/test?charset=utf8", 30)
	//绑定表
	orm.RegisterModel(new(UserInfo))
	db = orm.NewOrm()
}

func AddUser(user_info *UserInfo) (int64, error) {
	//插入
	return db.Insert(user_info)
}
