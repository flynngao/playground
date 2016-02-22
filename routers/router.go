package routers

import (
	"playground/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/",&controllers.MainController{})
  beego.Router("/run", &controllers.RunController{})

}