package models

import (
	"github.com/astaxie/beego/orm"
)

func Insert(model interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.Insert(model)
}
func InsertOneToMany(model interface{}, name string, sub ...interface{}) (int64, error) {
	orm := orm.NewOrm()
	m2m := orm.QueryM2M(model, name)
	return m2m.Add(sub...)
}
func Update(model interface{}, cols ...string) (int64, error) {
	orm := orm.NewOrm()
	return orm.Update(model, cols...)
}
func Delete(model interface{}) (int64, error) {
	orm := orm.NewOrm()
	return orm.Delete(model)
}

func GetModelWithConn(model interface{}, cols ...string) error {
	orm := orm.NewOrm()
	return orm.Read(model, cols...)
}
func QuertTableAll(ptrStructOrTableName string, model interface{}, filters map[string]interface{}) (int64, error) {
	qs := quertTableByFilter(ptrStructOrTableName, filters)
	return qs.All(model)
}
func QuertTableOne(ptrStructOrTableName string, model interface{}, filters map[string]interface{}) error {
	qs := quertTableByFilter(ptrStructOrTableName, filters)
	return qs.One(model)
}
func QuertTableAllRelatedSel(ptrStructOrTableName string, model interface{}, filters map[string]interface{}) orm.QuerySeter {
	return quertTableByFilter(ptrStructOrTableName, filters)
}

func quertTableByFilter(ptrStructOrTableName string, filters map[string]interface{}) orm.QuerySeter {
	orm := orm.NewOrm()
	qs := orm.QueryTable(ptrStructOrTableName)
	for key, val := range filters {
		qs = qs.Filter(key, val)
	}
	return qs
}

func LoadRelated(model interface{}, name string) (int64, error) {
	orm := orm.NewOrm()
	return orm.LoadRelated(model, name)
}

func init() {
	orm.Debug = true
}
