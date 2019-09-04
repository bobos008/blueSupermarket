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
	var (
		//var user User
		//var userSlice [] orm.ParamsList
		userMaps []orm.Params
	)
	o := orm.NewOrm()
	//_, err := o.QueryTable("user").ValuesList(&userSlice)
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

func (c *UserUpdateController) UserUpdate() {
	c.TplName = "blueTpl/userUpdate.html"
}

func (c *UserViewController) UserView() {
	c.TplName = "blueTpl/userView.html"
}
