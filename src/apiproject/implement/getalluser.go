package implement

import (
	"apiproject/models"
	"github.com/astaxie/beego/orm"
)

func GetAllUsers() (models.UserResponse, error) {
	var userResponse models.UserResponse
	var userList []models.UserInfo

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("userinfo")
	record, err := qs.All(&userList)

	userResponse.RecordNum = record
	userResponse.UserList = userList
	return userResponse, err
}