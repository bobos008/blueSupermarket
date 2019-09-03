package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type UserListController struct {
	beego.Controller
}

func (c *UserController) Index() {
	c.TplName = "blueTpl/index.tpl"
}

func (c *UserListController) UserList() {
	c.TplName = "blueTpl/userList.tpl"
}
