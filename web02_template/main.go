package main

import (
	"net/http"
	"text/template"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板
	//t, _ := template.ParseFiles("index.html")
	//通过MUST函数让Go自动处理异常
	t := template.Must(template.ParseFiles("index1.html", "index2.html"))
	//执行
	t.ExecuteTemplate(w, "index2.html", "hello index2.html")
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
