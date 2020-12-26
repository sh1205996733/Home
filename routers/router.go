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
	beego.Router("/api/v1.0/houses/index", &controllers.HomeIndexController{})

	/**用户注册*/
	beego.Router("/api/v1.0/users", &controllers.UserController{})
	/**用户登录*/
	beego.Router("/api/v1.0/sessions", &controllers.UserController{}, "post:Login")
	//api/v1.0/user/avatar 上传头像/更新用户名
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:UpdateAvatar")
	//api/v1.0/user 展示用户信息
	beego.Router("/api/v1.0/user", &controllers.UserController{}, "get:UserInfo")
	//api/v1.0/user/name  更新用户姓名
	beego.Router("/api/v1.0/user/name", &controllers.UserController{}, "put:UpdateName")
	//api/v1.0/user/auth  实名认证
	beego.Router("/api/v1.0/user/auth", &controllers.UserController{}, "get:UserInfo;post:Auth")

	//api/v1.0/user/houses 查看房源
	beego.Router("/api/v1.0/user/houses", &controllers.HouseController{}, "get:GetHouses;POST:PostHouse")
	//api/v1.0/houses 发布房源
	beego.Router("/api/v1.0/houses", &controllers.HouseController{}, "post:PostHouse")
	//api/v1.0/houses 上传房源图片
	beego.Router("/api/v1.0/houses/?:id/images", &controllers.HouseController{}, "post:HouseAvatar")
	//api/v1.0/houses/1
	beego.Router("/api/v1.0/houses/?:id", &controllers.HouseController{}, "get:HouseById")

	//api/v1.0/user/orders 订单
	beego.Router("/api/v1.0/user/orders", &controllers.OrderController{}, "get:GetOrderData")
}
