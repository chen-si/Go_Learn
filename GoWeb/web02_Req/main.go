package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "你发送的请求地址是：", r.URL.Path)
	// fmt.Fprintln(w, "你发送的请求后查询的字符串是：", r.URL.RawQuery)
	// fmt.Fprintln(w, "请求头中的所有信息：", r.Header)
	// fmt.Fprintln(w, "请求头中的Accept-Encoding信息：", r.Header["Accept-Encoding"])
	// fmt.Fprintln(w, "请求头中的Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))
	// //获取请求体内容的长度
	// length := r.ContentLength
	// //创建一个byte切片
	// body := make([]byte, length)
	// //将请求体的内容读取出来
	// r.Body.Read(body)

	//在浏览器中显示请求体的内容
	//fmt.Fprintln(w, "请求体中的内容：", string(body))

	//解析表单
	// r.ParseForm()

	// //如果form表单action属性的地址中也有与from表单参数名相同的参数
	// //那么参数值都可以得到，并且form表单中的参数在URL的前面
	// fmt.Fprintln(w, "请求参数有：", r.Form)
	// fmt.Fprintln(w, "请求参数有：", r.PostForm)

	//通过直接调用FormValue方法和PostFormValue直接获得请求参数
	fmt.Fprintln(w, "URL中的User请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "URL中的User请求参数的值是：", r.PostFormValue("username"))
}

func testJson(w http.ResponseWriter, r *http.Request) {
	//设置响应内容的类型
	w.Header().Set("Content-Type", "application/json")
	//创建User
	user := User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Email:    "liu@liu.com",
	}
	//将User转为json格式
	json, _ := json.Marshal(user)
	//将json格式数据响应给客户端
	w.Write(json)

}

func testRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJson)
	http.HandleFunc("/testRedirect", testRedirect)

	http.ListenAndServe(":8080", nil)
}
