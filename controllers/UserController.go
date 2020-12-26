package controllers

import (
	"Home/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tedcy/fdfs_client"
	"path"
)

type UserController struct {
	BaseController
}

func (this *UserController) Post() {
	resp := make(map[string]interface{})
	defer this.RespData(resp)

	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	user := models.User{}
	user.Password_hash = resp["password"].(string)
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)
	id, err := models.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_NODATA
		return
	}

	beego.Info("reg success ,id = ", id)
	resp["errno"] = models.RECODE_OK

	//放入session中
	this.SetSession("name", user.Name)
}

/**
登录
*/
func (this *UserController) Login() {

	//1.得到用户信息
	resp := make(map[string]interface{})
	defer this.RespData(resp)

	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	beego.Info("======name = ", resp["mobile"], "=======password =", resp["password"])
	//2.判断是否合法
	if resp["mobile"] == nil || resp["password"] == nil {
		resp["errno"] = models.RECODE_DATAERR
		return
	}

	//3.与数据库匹配判断账号密码正确

	user := models.User{Name: resp["mobile"].(string)}
	filterMap := make(map[string]interface{})
	filterMap["mobile__in"] = []string{user.Name, "xxxxxxxxxxxxxxxxxxx"}
	err := models.QuertTableOne("user", &user, filterMap)
	fmt.Println(user, "err -------'", err)
	if err != nil {
		resp["errno"] = models.RECODE_DATAERR
		return
	}

	if user.Password_hash != resp["password"] {
		resp["errno"] = models.RECODE_DATAERR
		beego.Info("333333name = ", resp["mobile"], "=======password =", resp["password"])
		return
	}

	//4.添加session
	this.SetSession("name", user.Name)
	this.SetSession("mobile", user.Mobile)
	this.SetSession("user_id", user.Id)

	//5.返回json数据给前端
	resp["errno"] = models.RECODE_OK
}

// 上传用户头像
func (this *UserController) UpdateAvatar() {
	resp := make(map[string]interface{})
	defer this.RespData(resp)
	//1.获取上传的一个文件
	fileData, hd, err := this.GetFile("avatar")
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

	//DataResponse.GroupName
	//DataResponse.RemoteFileId   //group/mm/00/00231312313131231.jpg

	//4.从session里拿到user_id
	user_id := this.GetSession("user_id")
	var user models.User
	//5.更新用户数据库中的内容
	filterMap := make(map[string]interface{})
	filterMap["Id"] = user_id
	models.QuertTableOne("user", &user, filterMap)
	user.Avatar_url = beego.AppConfig.String("imgUrl") + RemoteFileId

	_, errUpdate := models.Update(&user)
	if errUpdate != nil {
		resp["errno"] = models.RECODE_REQERR
		return
	}

	urlMap := make(map[string]string)
	urlMap["avatar_url"] = user.Avatar_url
	resp["errno"] = models.RECODE_OK
	resp["data"] = urlMap

}

// 获取个人信息
func (c *UserController) UserInfo() {
	resp := make(map[string]interface{})
	defer c.RespData(resp)

	// 获取用户id
	uid := c.GetSession("user_id")
	user := models.User{Id: uid.(int)}
	err := models.GetModelWithConn(&user)
	if err != nil {
		resp["errno"] = models.RECODE_REQERR
		return
	}
	resp["data"] = &user
	resp["errno"] = models.RECODE_OK
}

// 更新名字
func (c *UserController) UpdateName() {
	resp := make(map[string]interface{})
	defer c.RespData(resp)

	UserName := make(map[string]string)
	json.Unmarshal(c.Ctx.Input.RequestBody, &UserName)
	// 获取用户id
	uid := c.GetSession("user_id")
	user := models.User{Id: uid.(int)}
	if models.GetModelWithConn(&user) == nil {
		user.Name = UserName["name"]
		if _, err := models.Update(&user, "name"); err == nil {
			c.SetSession("name", UserName["name"])
			resp["data"] = UserName
			resp["errno"] = models.RECODE_OK
			return
		}
	}
	resp["errno"] = models.RECODE_DATAERR
}

//实名认证
func (this *UserController) Auth() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK

	defer this.RespData(resp)
	param := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &param)
	//获取Session中的userid
	userid := this.GetSession("user_id").(int)
	user := models.User{Id: userid}
	if models.GetModelWithConn(&user) == nil {
		user.Real_name = param["real_name"].(string)
		user.Id_card = param["id_card"].(string)
		if _, err := models.Update(&user); err == nil {
			resp["errno"] = models.RECODE_OK
			return
		}

	}
	resp["errno"] = models.RECODE_NODATA
}
