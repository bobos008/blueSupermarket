package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//var o orm.Ormer

func Syncdb() {
	createdb()
	Connect()
	//o := orm.NewOrm()
	// 数据库别名
	name := "default"
	// drop table 后在建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println("create tables error:", err)
	}
	insertUser()
	fmt.Println("database init is complete.")
}

//数据库连接
func Connect() {
	var dsn string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	// 连接mysql数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
	orm.RegisterDataBase("default", db_type, dsn)
}

// 创建数据库
func createdb() {
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	var dsn string
	var sqlstring string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
	sqlstring = fmt.Sprintf("CREATE DATABASE if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci",db_name)
	db, err := sql.Open(db_type, dsn)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, "created successful!")
	}
	defer db.Close()
}

// 添加超级管理员
func insertUser() {
	u := new(User)
	u.Username = "admin"
	u.Password = "admin"
	u.Createtime = time.Now()
	o := orm.NewOrm()
	o.Insert(u)
	fmt.Println("insert admin successful!")
}
