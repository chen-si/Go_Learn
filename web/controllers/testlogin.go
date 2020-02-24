package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type TestLoginController struct {
	beego.Controller
}

// type User struct {
// 	Username string
// 	PassWord string
// }

func (c *TestLoginController) Login() {
	name := c.Ctx.GetCookie("name")
	password := c.Ctx.GetCookie("password")

	// c.Ctx.SetCookie("name", name, -1, "/")清除cookie
	// c.Ctx.SetCookie("password", password, -1, "/")
	//c.Ctx.WriteString("Username " + name + " Password " + password)
	if name != "" {
		c.Ctx.WriteString("Username " + name + " Password " + password)
	} else {
		c.Ctx.WriteString(`<html><form action="http://localhost:8080/TestLogin" method="post">
								<input type="text" name="Username"/><br>
								<input type="password" name="PassWord"/><br>
								<input type="submit" value="提交"/><br>
							</form></html>`)
	}
}

func (c *TestLoginController) Post() {
	user := &User{}
	if err := c.ParseForm(user); err != nil {
		fmt.Println(err)
		return
	}
	c.SetSession("name", user.Username)
	c.SetSession("password", user.PassWord)
	c.Ctx.SetCookie("name", user.Username, 100, "/")
	c.Ctx.SetCookie("password", user.PassWord, 100, "/")
	c.Ctx.WriteString("Username" + user.Username + " Password" + user.PassWord)
}
