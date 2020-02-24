package main

import (
	"GoWeb/web03_action/model"
	"html/template"
	"net/http"
)

//测试 If
func testIf(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	//执行
	t.Execute(w, age < 18)
}

//测试 Range
func testRange(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("range.html"))
	var emps []*model.Emlpoyee
	emp := &model.Emlpoyee{
		ID:       1,
		LastName: "liu",
		Email:    "liu@liu.com",
	}
	emps = append(emps, emp)
	emp = &model.Emlpoyee{
		ID:       2,
		LastName: "yang",
		Email:    "yang@liu.com",
	}
	emps = append(emps, emp)
	emp = &model.Emlpoyee{
		ID:       3,
		LastName: "wang",
		Email:    "wang@liu.com",
	}
	emps = append(emps, emp)
	//执行
	t.Execute(w, emps)
}

func testWith(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("with.html"))
	//执行
	t.Execute(w, "狸猫")
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	//执行
	t.Execute(w, "我 能在两个文件中显示吗")
}

func testDefine(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("define.html"))
	//执行
	t.ExecuteTemplate(w, "model", "")
}

func testDefine2(w http.ResponseWriter, r *http.Request) {
	age := -25
	var t *template.Template
	if age < 0 {
		t = template.Must(template.ParseFiles("define2.html"))
	} else if age < 18 {
		t = template.Must(template.ParseFiles("define2.html", "content1.html"))
	} else {
		t = template.Must(template.ParseFiles("define2.html", "content2.html"))
	}

	//执行
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/testIf", testIf)

	http.HandleFunc("/testRange", testRange)

	http.HandleFunc("/testWith", testWith)

	http.HandleFunc("/testTemplate", testTemplate)

	http.HandleFunc("/testDefine", testDefine)

	http.HandleFunc("/testDefine2", testDefine2)

	http.ListenAndServe(":8080", nil)
}
