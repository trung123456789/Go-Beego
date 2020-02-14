package implement

import (
	"apiproject/models"
	"github.com/astaxie/beego/orm"
)

func GetUser(u models.UserInfo) (models.UserResponse, error) {
	var userList []models.UserInfo
	var userResponse models.UserResponse
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("userinfo")
	// Get a QuerySeter object. User is table name

	if u.Name != "" {
		qs = qs.Filter("name__icontains", u.Name)
	}
	if u.UserId != "" {
		qs = qs.Filter("user_id", u.UserId)
	}
	if u.Age != 0 {
		qs = qs.Filter("age", u.Age)
	}

	num, err := qs.All(&userList)
	userResponse.RecordNum = num
	userResponse.UserList = userList

	return userResponse, err

}