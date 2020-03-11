package main

import (
	"fmt"
	"net/http"
)

//setCookie 添加cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	//创建cookie
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
	}
	// //发送给浏览器
	// w.Header().Set("Set-Cookie", cookie.String())
	// //添加第二个cookie
	// w.Header().Add("Set-Cookie", cookie.String())

	//直接调用SetCookie
	http.SetCookie(w, &cookie)

}

//getCookies
func getCookies(w http.ResponseWriter, r *http.Request) {
	//获取请求头中的所有cookie
	cookies := r.Header["Cookie"]
	cookie, _ := r.Cookie("user")
	fmt.Println("得到的Cookie：", cookies, cookie)
}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookies)

	http.ListenAndServe(":8080", nil)
}
