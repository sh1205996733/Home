package controllers

import (
	"Home/models"
)

type SessionController struct {
	BaseController
}

/**
获取用户Session
*/
func (this *SessionController) Get() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_DBERR
	defer this.RespData(resp)
	user := models.User{}
	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["data"] = user
	}
}

/**
退出登录
*/
func (this *SessionController) Delete() {
	resp := make(map[string]interface{})
	defer this.RespData(resp)
	this.DelSession("name")
	resp["errno"] = models.RECODE_OK
}
