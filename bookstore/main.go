package main

import (
	"net/http"
	"text/template"
)


//去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("index.html"))

	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/main",IndexHandler)

	http.ListenAndServe(":8080",nil)
}
