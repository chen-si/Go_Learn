package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type TestInputController struct {
	beego.Controller
}

type User struct {
	Username string
	PassWord string
}

func (c *TestInputController) Get() {
	// id := c.GetString("id")
	// c.Ctx.WriteString(id)

	name := c.GetSession("name")
	password := c.GetSession("password")

	if name != ""{
		c.Ctx.WriteString(name.(string)+" "+password.(string))
	}else{
		c.Ctx.WriteString(`<html><form action="http://localhost:8080/TestInput" method="post">
								<input type="text" name="Username"/><br>
								<input type="password" name="PassWord"/><br>
								<input type="submit" value="提交"/><br>
							</form></html>`)
	}

	// name := c.Input().Get("name")
	// 
	
}

func (c *TestInputController) Post() {
	user := &User{}
	if err := c.ParseForm(user); err != nil {
		fmt.Println(err)
		return
	}
	c.Ctx.WriteString("Username" + user.Username + " Password" + user.PassWord)

}
