package controllers

import (
	"Home/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tedcy/fdfs_client"
	"path"
	"strconv"
)

/**
房源
*/
type HouseController struct {
	BaseController
}

/**
获取房源
*/
func (this *HouseController) GetHouses() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	defer this.RespData(resp)
	//1.从Session中获取user_id
	user_id := this.GetSession("user_id").(int)
	//2.查询数据库
	houses := []models.House{}
	//num, err := o.QueryTable("post").Filter("User", 1).RelatedSel().All(&posts)
	filter := make(map[string]interface{})
	filter["user_id"] = user_id
	num, err := models.QuertTableAll("house", &houses, filter)
	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
	}
	if num == 0 {
		resp["essno"] = models.RECODE_NODATA
	}
	fmt.Println("user----------", houses)
	//3.返回前端
	respData := make(map[string]interface{})
	respData["houses"] = &houses
	resp["data"] = &respData
	return
}

/**
发布房源
*/
func (this *HouseController) PostHouse() {
	resp := make(map[string]interface{})
	defer this.RespData(resp)

	//1 获取前端数据
	params := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &params)

	//2. 插入数据到数据库中
	//房间信息
	house := models.House{}
	house.Title = params["title"].(string)
	price, _ := strconv.Atoi(params["price"].(string))
	house.Price = price
	house.Address = params["address"].(string)
	room_count, _ := strconv.Atoi(params["room_count"].(string))
	house.Room_count = room_count
	house.Unit = params["unit"].(string)
	house.Beds = params["beds"].(string)
	min_days, _ := strconv.Atoi(params["min_days"].(string))
	house.Min_days = min_days
	max_days, _ := strconv.Atoi(params["max_days"].(string))
	house.Max_days = max_days

	// 设备信
	facility := []*models.Facility{}
	for _, fid := range params["facility"].([]interface{}) {
		beego.Info("fid:", fid)
		f_id, _ := strconv.Atoi(fid.(string))
		fac := &models.Facility{Id: f_id}
		facility = append(facility, fac)
	}

	area_id, _ := strconv.Atoi(params["area_id"].(string))
	area := models.Area{Id: area_id}
	house.Area = &area

	// t填充用户信息
	uid := this.GetSession("user_id")
	user := models.User{Id: uid.(int)}
	house.User = &user

	house_id, err := models.Insert(&house)
	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		return
	}
	house.Id = int(house_id)
	num, err := models.InsertOneToMany(&house, "Facilities", facility)
	if err != nil || num == 0 {
		resp["errno"] = models.RECODE_NODATA
		return
	}
	respData := make(map[string]interface{})
	respData["house_id"] = &house_id
	resp["data"] = &respData
	resp["errno"] = models.RECODE_OK
}

/**
上传房源图片
*/
func (this *HouseController) HouseAvatar() {
	resp := make(map[string]interface{})
	defer this.RespData(resp)
	//1.获取上传的一个文件
	fileData, hd, err := this.GetFile("house_image")
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		return
	}
	//2.得到文件后缀
	suffix := path.Ext(hd.Filename) //a.jpg.avi

	//3.存储文件到fastdfs上
	fdfsClient, err := fdfs_client.NewClientWithConfig("conf/fdfs.conf")
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		return
	}
	fileBuffer := make([]byte, hd.Size)
	_, err = fileData.Read(fileBuffer)
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		return
	}
	RemoteFileId, err := fdfsClient.UploadByBuffer(fileBuffer, suffix[1:]) //aa.jpg

	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		return
	}

	//4.从前端拿到house_id
	house_id := this.GetString("house_id")
	var house models.House
	//5.更新house数据库中的内容
	filterMap := make(map[string]interface{})
	filterMap["Id"] = house_id
	models.QuertTableOne("house", &house, filterMap)
	house.Index_image_url = beego.AppConfig.String("imgUrl") + RemoteFileId

	_, errUpdate := models.Update(&house)
	if errUpdate != nil {
		resp["errno"] = models.RECODE_REQERR
		return
	}

	urlMap := make(map[string]string)
	urlMap["index_image_url"] = house.Index_image_url
	resp["errno"] = models.RECODE_OK
	resp["data"] = urlMap
}

/**
查看房源详情
*/
func (this *HouseController) HouseById() {
	resp := make(map[string]interface{})
	defer this.RespData(resp)

	uid := this.GetSession("user_id")
	user := models.User{Id: uid.(int)}
	house_id := this.Ctx.Input.Param(":id")
	hid, _ := strconv.Atoi(house_id)

	house := models.House{Id: hid}
	house.User = &user

	if err := models.GetModelWithConn(&house); err != nil {
		resp["errno"] = models.RECODE_DATAERR
		return
	}
	// 载入关联字段
	models.LoadRelated(&house, "Area")
	models.LoadRelated(&house, "User")
	models.LoadRelated(&house, "Images")
	models.LoadRelated(&house, "Facilities")

	facs := []string{}
	for _, fac := range house.Facilities {
		fid := strconv.Itoa(fac.Id)
		facs = append(facs, fid)
	}
	respData := make(map[string]interface{})
	respData["facilities"] = facs
	respData["house"] = house

	resp["data"] = respData
	resp["errno"] = models.RECODE_OK
}
