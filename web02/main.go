package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器处理请求")
}

func main() {
	myHandler := &MyHandler{}

	//http.Handle("/myHandler", myHandler)
	server := http.Server{
		Addr:        ":8848",
		Handler:     myHandler,
		ReadTimeout: 2 * time.Second,
	}
	err := server.ListenAndServe()
	fmt.Println(err)
}
