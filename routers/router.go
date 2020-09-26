package routers

import (
	"BeegoProject001/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.MainController{})
    beego.Router("/regidter",&controllers.RegidterController{})
}
