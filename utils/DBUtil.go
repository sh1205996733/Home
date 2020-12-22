package utils

import (
	. "Home/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	driverName := beego.AppConfig.String("driverName")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	// set default database
	dbConn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	beego.Info("dbConn = ", dbConn)
	orm.RegisterDataBase("default", driverName, dbConn)
	orm.RegisterModel(new(User), new(OrderHouse), new(Facility), new(HouseImage), new(Area), new(House))
	orm.RunSyncdb("default", false, true)
}
