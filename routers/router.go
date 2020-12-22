package routers

import (
	"Home/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{})
	beego.Router("/api/v1.0/session", &controllers.SessionController{})
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{})
	beego.Router("/api/v1.0/users", &controllers.UserController{})
}
