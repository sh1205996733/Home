package controllers

import (
	"Home/models"
)

type HomeIndexController struct {
	BaseController
}

func (this *HomeIndexController) Get() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_DBERR
	this.RespData(resp)
}
