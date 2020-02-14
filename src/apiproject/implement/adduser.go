package implement

import (
	"apiproject/conf"
	"apiproject/models"
	"errors"
	"github.com/astaxie/beego/orm"
)

func AddUser(u models.UserInfo) (string, error) {
	if u.UserId == "" || u.Password == "" {
		return "", errors.New(conf.IniConf.String("missingIdOrPass"))
	}
	o := orm.NewOrm()
	o.Using("default")

	user := models.UserInfo{
		UserId:    u.UserId,
		Name:      u.Name,
		Age:       u.Age,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	_, err := o.Insert(&user)
	return u.UserId, err
}
