package routers

import (
	"github.com/astaxie/beego"
	"github.com/guanzi008/go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
