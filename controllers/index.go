package controllers

import(
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

type PasswordController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	c.TplName = "blueTpl/index.html"
}

func (c *LoginController) Login() {
	c.TplName = "blueTpl/login.html"
}

func (c *PasswordController) Password() {
	c.TplName = "blueTpl/password.html"
}
