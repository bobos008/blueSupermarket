package main

import (
	"fmt"
	"os"
	_ "blueSupermarket/routers"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("Starting......")
	fmt.Println("Start ok")
	initialize()
	beego.Run()
}

func initialize() {
	initArgs()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		fmt.Println(v)
		/*if v == "-syncdb" {
		}*/
	}
}
