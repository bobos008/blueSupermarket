package controllers

import(
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

type CheckLoginController struct {
	beego.Controller
}

type LogoutController struct {
	beego.Controller
}

type PasswordController struct {
	beego.Controller
}

type UpdatePasswordController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	c.TplName = "blueTpl/index.html"
}


func (c *LoginController) Login() {
	ssUser := c.GetSession("user")
	if ssUser != nil {
		c.Redirect("/", 302)
	}
	c.TplName = "blueTpl/login.html"
}

func (c *LogoutController) Logout() {
	c.DelSession("user")
	c.TplName = "blueTpl/login.html"
}

func (c *CheckLoginController) CheckLogin() {
	user := c.GetString("username")
	password := c.GetString("password")
	o := orm.NewOrm()
	isUser := o.QueryTable("user").Filter("Username", user).Filter("Password", password).Exist()
	isSuccess := false
	if isUser {
		isSuccess = true
		c.SetSession("user", user)
	}
	c.Data["json"] = isSuccess
	c.ServeJSON()
}

func (c *PasswordController) Password() {
	ssUser := c.GetSession("user")
	if ssUser == nil {
		c.Redirect("/login", 302)
	}
	c.TplName = "blueTpl/password.html"
}

func (c *UpdatePasswordController) UpdatePassword() {
	ssUser := c.GetSession("user")
	if ssUser == nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}
	oldPassword := c.GetString("oldPassword")
	newPassword := c.GetString("newPassword")
	o := orm.NewOrm()
	_, err := o.QueryTable("user").Filter("Username", ssUser).Filter("Password",oldPassword).Update(orm.Params{
		"password": newPassword,
	})
	if err == nil {
		c.Data["json"] = true
		c.ServeJSON()
		return
	}
	c.Data["json"] = false
	c.ServeJSON()
}
