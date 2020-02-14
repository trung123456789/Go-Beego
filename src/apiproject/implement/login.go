package implement

import (
	"apiproject/models"
	"github.com/astaxie/beego/orm"
	"reflect"
)

func Login(id, password string) (bool, error) {

	o := orm.NewOrm()
	o.Using("default")

	// Get a QuerySeter object. User is table name
	qs := o.QueryTable("userinfo")
	userInfo := models.UserInfo{}
	errSelect := qs.Filter("UserId", id).Filter("Password", password).One(&userInfo)
	if errSelect != nil {
		return false, errSelect
	}
	if !reflect.DeepEqual(userInfo, models.UserInfo{}) {
		return true, nil
	}
	return false, nil
}
