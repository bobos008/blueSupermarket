package controllers

import (
	. "blueSupermarket/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BillListController struct {
	beego.Controller
}

type BillAddController struct {
	beego.Controller
}

type BillAddDataController struct {
	beego.Controller
}

func (c *BillListController) BillList() {
	c.TplName = "blueTpl/billList.html"
}

func (c *BillAddController) BillAdd() {
	var providerMaps []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable("provider").Values(&providerMaps, "Id", "ProviderName")
	if err == nil {
		c.Data["provider"] = providerMaps
	} else {
		c.Data["provider"] = nil
	}
	c.TplName = "blueTpl/billAdd.html"
}

func (c *BillAddDataController) BillAddData() {
	billId := c.GetString("billId")
	billName := c.GetString("billName")
	billCom := c.GetString("billCom")
	billNum := c.GetString("billNum")
	money := c.GetString("money")
	supplier := c.GetString("supplier")
	zhifu := c.GetString("zhifu")

	var bill Bill
	bill.GoodsName = billName
	bill.GoodsNumber = billId

	c.Data["json"] = true
	c.ServeJSON()
}
