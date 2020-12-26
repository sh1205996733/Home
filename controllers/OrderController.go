package controllers

import (
	"Home/models"
)

/**
订单
*/
type OrderController struct {
	BaseController
}

// 订单
func (this *OrderController) GetOrderData() {
	resp := make(map[string]interface{})
	defer this.RespData(resp)

	uid := this.GetSession("user_id").(int)

	// 根据url获取当前操作的角色
	role := this.GetString("role")
	if role == "landlord" { //custom
		orders := []models.OrderHouse{}
		user := models.User{Id: uid}
		filter := make(map[string]interface{})
		filter["user_id"] = uid
		models.QuertTableAll("OrderHouse", &orders, filter)
		for _, order := range orders {
			order.User = &user
			models.LoadRelated(order, "User")
		}

		respData := make(map[string]interface{})
		respData["orders"] = orders
		resp["data"] = respData
		resp["errno"] = models.RECODE_OK
		return
	}

	if role == "landlord" {

	}

	if role == "" {
		resp["errno"] = models.RECODE_ROLEERR
		return
	}
}
