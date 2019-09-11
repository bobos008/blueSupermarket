package controllers

import (
	. "blueSupermarket/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
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
	goodsCount,err := strconv.ParseUint(billNum, 10, 32)
	if err != nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}
	money := c.GetString("money")
	moneyFloat64, err := strconv.ParseFloat(money, 64)
	if err != nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}
	account, err := strconv.ParseFloat(fmt.Sprintf("%.2f", moneyFloat64),32)
	if err != nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}
	supplier := c.GetString("supplier")
	providerId, err := strconv.ParseInt(supplier,10,64)
	if err != nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}
	zhiFu := c.GetString("zhifu")
	isPay,err := strconv.ParseBool(zhiFu)
	if err != nil {
		c.Data["json"] = false
		c.ServeJSON()
		return
	}

	var bill Bill
	bill.GoodsName = billName
	bill.GoodsNumber = billId
	bill.GoodsUnit = billCom
	bill.GoodsCount = uint32(goodsCount)
	bill.Amount = account
	bill.IsPay = isPay
	bill.Provider = &Provider{Id:providerId}

	o := orm.NewOrm()
	id, err := o.Insert(&bill)
	if id == 0 {
		fmt.Println(err)
		c.Data["json"] = false
	} else {
		c.Data["json"] = true
	}
	c.ServeJSON()
}
