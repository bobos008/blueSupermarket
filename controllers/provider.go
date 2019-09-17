package controllers

import (
	. "blueSupermarket/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type ProviderController struct {
	beego.Controller
}

type ProviderAddController struct {
	beego.Controller
}

type ProviderAddDataController struct {
	beego.Controller
}

type ProviderDelController struct {
	beego.Controller
}

type ProviderViewController struct {
	beego.Controller
}

type ProviderUpdateController struct {
	beego.Controller
}

type ProviderUpdateDataController struct {
	beego.Controller
}

func (c *ProviderController) ProviderList() {
	var providerMaps []orm.Params
	o := orm.NewOrm()
	providerName := c.GetString("providerName")
	_, err := o.QueryTable("provider").Filter("providername__icontains", providerName).OrderBy("-id").Values(&providerMaps)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, providerMap := range providerMaps {
			createTime := providerMap["Createtime"].(time.Time)
			createTimeFormat := time.Unix(createTime.Unix(), 0).Format("2006-01-02")
			providerMap["Createtime"] = createTimeFormat
		}
	}
	c.Data["provider"] = &providerMaps
	c.TplName = "blueTpl/providerList.html"
}

func (c *ProviderAddController) ProviderAdd() {
	c.TplName = "blueTpl/providerAdd.html"
}

func (c *ProviderAddDataController) ProviderAddData() {
	providerId := c.GetString("providerId")
	providerName := c.GetString("providerName")
	people := c.GetString("people")
	phone := c.GetString("phone")
	address := c.GetString("address")
	fax := c.GetString("fax")
	describe := c.GetString("describe")

	var provider Provider
	provider.ProviderNumber = providerId
	provider.ProviderName = providerName
	provider.People = people
	provider.PhoneNumber = phone
	provider.Address = address
	provider.Fax = fax
	provider.Describe = describe

	o := orm.NewOrm()
	id, err := o.Insert(&provider)
	isSuccess := true
	if id == 0 {
		fmt.Println(err)
		isSuccess = false
	}

	c.Data["json"] = isSuccess
	c.ServeJSON()
}

func (c *ProviderDelController) ProviderDel() {
	id := c.GetString("id")
	intId, err := strconv.ParseInt(id,10,64)
	if err != nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	if _, err:=o.Delete(&Provider{Id:intId});err == nil {
		c.Data["json"] = true
		c.ServeJSON()
		return
	}
	c.Data["json"] = false
	c.ServeJSON()
}

func (c *ProviderViewController) ProviderView() {
	var providerMaps []orm.Params
	id := c.GetString("id")
	o := orm.NewOrm()
	_, err := o.QueryTable("provider").Filter("Id",id).Limit(1).Values(&providerMaps)
	if err != nil {
		fmt.Println(err)
		c.Data["provider"] = nil
	} else {
		c.Data["provider"] = providerMaps
	}
	c.TplName = "blueTpl/providerView.html"
}

func (c *ProviderUpdateController) ProviderUpdate() {
	var providerMaps []orm.Params
	id := c.GetString("id")
	o := orm.NewOrm()
	_,err := o.QueryTable("provider").Filter("Id",id).Values(&providerMaps)
	if err != nil {
		fmt.Println(err)
		c.Data["provider"] = nil
	} else {
		c.Data["provider"] = providerMaps
	}
	c.TplName = "blueTpl/providerUpdate.html"
}

func (c *ProviderUpdateDataController) ProviderUpdateData() {
	id := c.GetString("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}
	providerId := c.GetString("providerId")
	providerName := c.GetString("providerName")
	people := c.GetString("people")
	phone := c.GetString("phone")
	address := c.GetString("address")
	fax := c.GetString("fax")
	describe := c.GetString("describe")

	o := orm.NewOrm()
	provider := Provider{Id:intId}
	if o.Read(&provider) == nil {
		provider.ProviderNumber = providerId
		provider.ProviderName = providerName
		provider.People = people
		provider.PhoneNumber = phone
		provider.Address = address
		provider.Fax = fax
		provider.Describe = describe
		_,err := o.Update(&provider)
		if err == nil {
			c.Data["json"] = true
			c.ServeJSON()
			return
		}
	}
	c.Data["json"] = false
	c.ServeJSON()
}
