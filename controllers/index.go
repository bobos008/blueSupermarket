package controllers

import (
	"fmt"
	"time"
	"strconv"
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
	c.TplName = "blueTpl/userList.html"
}

func (c *UserAddController) UserAdd() {
	c.TplName = "blueTpl/userAdd.html"
}

func (c *UserAddDataController) UserAddData() {
	is_success := true
	username := c.GetString("userName")
	fmt.Println(username)
	password := c.GetString("password")
	gender := c.GetString("gender")
	phone := c.GetString("phone")
	birthday := c.GetString("birthday")
	/*loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", birthday + " 00:00:00", loc)
	var unix_time int64
	if err == nil {
		unix_time = the_time.Unix()
	} else {
		fmt.Println(err)
		unix_time = time.Now().Unix()
	}
	time_date := time.Unix(unix_time, 0)*/
	time_date,err := time.Parse("2006-01-02 15:04:05", birthday + " 00:00:00")
	if err != nil {
		fmt.Println(err)
		time_date = time.Now()
	}

	res_gender := false
	if gender == "man" {
		res_gender = true
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
	user.Gender = res_gender
	user.Phone_number = phone
	user.Birthday = time_date
	user.Site = address
	user.Role = role

	id, err:= o.Insert(&user)
	if id == 0{
		fmt.Println(err)
		is_success = false
	}
	c.Data["json"] = is_success
	c.ServeJSON()
}

func (c *UserUpdateController) UserUpdate() {
	c.TplName = "blueTpl/userUpdate.html"
}

func (c *UserViewController) UserView() {
	c.TplName = "blueTpl/userView.html"
}
