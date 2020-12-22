package controllers

import (
	"Home/models"
	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RespData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *SessionController) Get() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	defer this.RespData(resp)
	user := models.User{}
	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}
}
func (this *SessionController) Delete() {
	resp := make(map[string]interface{})
	defer this.RespData(resp)
	this.DelSession("name")
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
