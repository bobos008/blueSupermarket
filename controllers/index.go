package controllers

import (
	"fmt"
	"strconv"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "blueSupermarket/models"
)

type UserController struct {
	beego.Controller
}

type UserListController struct {
	beego.Controller
}

type UserAddController struct {
	beego.Controller
}

type UserAddDataController struct {
	beego.Controller
}

type UserDelController struct {
	beego.Controller
}

type UserUpdateController struct {
	beego.Controller
}

type UserViewController struct {
	beego.Controller
}

func (c *UserController) Index() {
	c.TplName = "blueTpl/index.html"
}

func (c *UserListController) UserList() {
	var userMaps []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable("user").OrderBy("-id").Values(&userMaps)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, userMap := range userMaps {
			birthday := userMap["Birthday"]
			if birthday != nil {
				bornYear := birthday.(time.Time).Year()
				currentYear := time.Now().Year()
				userMap["Age"] = currentYear - bornYear
			} else {
				userMap["Age"] = 0
			}
		}
	}
	c.Data["user"] = &userMaps
	c.TplName = "blueTpl/userList.html"
}

func (c *UserAddController) UserAdd() {
	c.TplName = "blueTpl/userAdd.html"
}

func (c *UserAddDataController) UserAddData() {
	isSuccess := true
	username := c.GetString("userName")
	fmt.Println(username)
	password := c.GetString("password")
	gender := c.GetString("gender")
	phone := c.GetString("phone")
	birthday := c.GetString("birthday")
	timeDate,err := time.Parse("2006-01-02", birthday)
	if err != nil {
		fmt.Println(err)
		timeDate = time.Now()
	}

	resGender := false
	if gender == "man" {
		resGender = true
	}
	address := c.GetString("address")
	userLei := c.GetString("userlei")
	role,err := strconv.Atoi(userLei)
	if err != nil {
		role = 3
	}

	o := orm.NewOrm()
	var user User
	user.Username = username
	user.Password = password
	user.Gender = resGender
	user.Phone_number = phone
	user.Birthday = timeDate
	user.Site = address
	user.Role = role

	id, err:= o.Insert(&user)
	if id == 0{
		fmt.Println(err)
		isSuccess = false
	}
	c.Data["json"] = isSuccess
	c.ServeJSON()
}

func (c *UserDelController) UserDel() {
	id := c.GetString("id")
	intId,err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	if _,err := o.Delete(&User{Id:intId});err == nil {
		c.Data["json"] = true
		c.ServeJSON()
		return
	}
	c.Data["json"] = false
	c.ServeJSON()
}

func (c *UserUpdateController) UserUpdate() {
	c.TplName = "blueTpl/userUpdate.html"
}

func (c *UserViewController) UserView() {
	var userMaps []orm.Params
	o := orm.NewOrm()
	id := c.GetString("id")
	_, err := o.QueryTable("user").Filter("Id", id).Limit(1).Values(&userMaps)
	if err != nil {
		fmt.Println(err)
		c.Data["user"] = nil
	} else {
		c.Data["user"] = userMaps
	}
	c.TplName = "blueTpl/userView.html"
}
