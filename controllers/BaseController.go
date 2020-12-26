package controllers

import (
	"Home/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) RespData(resp map[string]interface{}) {
	resp["errmsg"] = models.RecodeText(resp["errno"].(string))
	this.Data["json"] = resp
	this.ServeJSON()
}
