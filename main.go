package main

import (
	"blueSupermarket/models"
	_ "blueSupermarket/routers"
	"fmt"
	//"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"os"
)

func main() {
	fmt.Println("Starting......")
	fmt.Println("Start ok")
	initialize()
	beego.BConfig.WebConfig.Session.SessionOn = true

	// 过滤器(验证是否登录)
	var filterUser = func(ctx *context.Context) {
		const loginUrl string = "/login"
		user := ctx.Input.Session("user")
		isAjax := ctx.Input.IsAjax()
		RqUrl := ctx.Request.RequestURI

		if(RqUrl != loginUrl){
			if(user == nil) && (!isAjax){
				ctx.Redirect(302, "/login")
			}
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, filterUser)

	beego.Run()
}

func initialize() {
	initArgs()
	models.Connect()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		fmt.Println(v)
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}
