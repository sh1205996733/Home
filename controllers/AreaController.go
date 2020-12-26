package controllers

import (
	"Home/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

type AreaController struct {
	BaseController
}

func (this *AreaController) Get() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	defer this.RespData(resp)

	//从redis缓存中拿数据拿数据
	cache_conn, err := cache.NewCache("redis", beego.AppConfig.String("redis"))
	if areaData := cache_conn.Get("area"); areaData != nil {
		beego.Info("get data from cache===========")
		area := []models.Area{}
		err := json.Unmarshal(areaData.([]byte), &area)
		if err != nil {
			resp["errno"] = models.RECODE_DBERR
			return
		}
		resp["data"] = area
		return
	}

	//从mysql数据库拿到area数据
	var areas []models.Area
	num, err := models.QuertTableAll("area", &areas, nil)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		return
	}
	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		return
	}
	resp["data"] = areas

	//把数据转换成json格式存入缓存
	json_str, err := json.Marshal(areas)
	if err != nil {
		beego.Info("encoding err")
		return
	}
	cache_conn.Put("area", json_str, time.Second*3600)
	//打包成json返回给前段
	beego.Info("query data sucess ,resp =", resp, "num =", num)
}
