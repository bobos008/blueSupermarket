package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 供应商表
type Provider struct {
	Id int64
	ProviderNumber string `orm:"unique;size(32)"`
	ProviderName string `orm:"size(32)"`
	People string `orm:"size(32)"`
	PhoneNumber string `orm:"size(11)"`
	Address string `orm:"size(500);null"`
	Fax string `orm:"size(11);null"`
	Describe string `orm:"size(500);null"`
	Createtime time.Time `orm:"type(datetime);auto_now_add"`
}

func init() {
	orm.RegisterModel(new(Provider))
}
