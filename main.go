package main

import (
	_ "Home/routers"
	_ "Home/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

func main() {
	//utils.TestUploadByFilename("main.go")
	ignoreStaticPath()
	beego.Run()
	//beego.Run(":8088")
}

/**
配置前端过滤器
*/
func ignoreStaticPath() {
	//透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}
func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
