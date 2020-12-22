package controllers

import (
	"Home/models"
	"github.com/astaxie/beego"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController) RespData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *HouseIndexController) Get() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	this.RespData(resp)
}
