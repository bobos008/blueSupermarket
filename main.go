package main

import (
	"blueSupermarket/models"
	_ "blueSupermarket/routers"
	"fmt"
	"github.com/astaxie/beego"
	"os"
)

func main() {
	fmt.Println("Starting......")
	fmt.Println("Start ok")
	initialize()
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
