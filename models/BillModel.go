package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

// 账单管理
type Bill struct {
	Id int64
	GoodsNumber string `orm:"size(60)"`
	GoodsName string `orm:"size(50)"`
	GoodsUnit string `orm:"size(20)"`
	GoodsCount uint32
	Amount float64 `orm:"digits(12);decimals(2)"`
	IsPay bool `orm:"default(false)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	Provider *Provider `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Bill))
}
