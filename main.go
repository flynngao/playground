package main

import (
	_ "playground/routers"
	"github.com/astaxie/beego"
)

func main() {
  beego.SetStaticPath("/static","static")
	beego.Run()
}

