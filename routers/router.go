package routers

import (
	"Home/controllers"
	"github.com/astaxie/beego"
)

func init() {
	/**地区列表*/
	beego.Router("/api/v1.0/areas", &controllers.AreaController{})
	/**Session处理*/
	beego.Router("/api/v1.0/session", &controllers.SessionController{})
	/**房屋首页*/
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{})
	/**用户登录*/
	beego.Router("/api/v1.0/users", &controllers.UserController{})
}
