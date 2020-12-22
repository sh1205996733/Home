package models

import "github.com/astaxie/beego/orm"

func Insert(model interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.Insert(model)
}
func Update(model interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.Update(model)
}
func Delete(model interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.Delete(model)
}
func GetUserWithConn(model interface{}, conn ...string) error {
	orm := orm.NewOrm()
	return orm.Read(model, conn...)
}
func QuertTableAll(ptrStructOrTableName string, model interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.QueryTable(ptrStructOrTableName).All(model)
}
