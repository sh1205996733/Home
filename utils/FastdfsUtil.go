package utils

import (
	"github.com/astaxie/beego"
	"github.com/tedcy/fdfs_client"
)

func TestUploadByFilename(fileName string) (groupName string, fileId string, err error) {
	fdfsClient, errClient := fdfs_client.NewClientWithConfig("conf/fdfs.conf")
	if errClient != nil {
		beego.Info("New FdfsClient error %s", errClient.Error())
		return "", "", errClient
	}
	fileId, errUpload := fdfsClient.UploadByFilename(fileName)
	if errUpload != nil {
		beego.Info("New FdfsClient error %s", errUpload.Error())
		return "", "", errUpload
	}
	beego.Info("=================groupNmae = ", "")
	beego.Info("=================fileId = ", fileId)
	return "", fileId, nil
}
