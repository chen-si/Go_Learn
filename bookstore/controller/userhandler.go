package controller

import (
	"bookstore/dao"
	"html/template"
	"net/http"
)

//Login处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用验证用户名密码的方法
	user, _ := dao.CheckUsernameAndPassword(username, password)
	if user.ID > 0 {
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		//用户名和密码不正确
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	user, _ := dao.CheckUsername(username)
	if user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		//用户名可以用
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}

}

//CheckUserName 通过发送Ajax请求验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")

	user, _ := dao.CheckUsername(username)
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在!"))
	} else {
		//用户名可以用
		w.Write([]byte("<font style='color:green'>用户名可用</font>"))
	}
}
