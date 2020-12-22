package controllers

import (
	"Home/models"
	"github.com/astaxie/beego"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RespData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *AreaController) Get() {
	beego.Info("connect success")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RespData(resp)

	//从mysql数据库拿到area数据
	var areas []models.Area
	num, err := models.QuertTableAll("area", &areas)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	resp["data"] = areas
}
