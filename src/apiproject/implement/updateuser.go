package implement

import (
	"apiproject/conf"
	"apiproject/models"
	"errors"
	"github.com/astaxie/beego/orm"
)

func UpdateUser(uid string, uu models.UserInfo) (map[string]string, error) {
	m := make(map[string]string)
	var err error
	user := models.UserInfo{
		UserId: uid,
	}
	o := orm.NewOrm()
	o.Using("default")

	if err = o.Read(&user, "UserId"); err == nil {
		if uu.Name != "" {
			user.Name = uu.Name
		}
		if uu.Age != 0 {
			user.Age = uu.Age
		}
		if num, errUpdate := o.Update(&user); errUpdate == nil && num > 0 {
			m["user_id"] = user.UserId
			return m, errUpdate
		}
	} else {
		err = errors.New(conf.IniConf.String("userNotExist"))
	}

	return m, err
}
