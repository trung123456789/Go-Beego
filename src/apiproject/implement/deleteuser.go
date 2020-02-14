package implement

import (
	"apiproject/conf"
	"apiproject/models"
	"errors"
	"github.com/astaxie/beego/orm"
)

func DeleteUser(uid string) (map[string]string, error) {
	var err error
	var num int64
	m := make(map[string]string)
	o := orm.NewOrm()
	if num, err = o.Delete(&models.UserInfo{UserId: uid}, "UserId"); err == nil && num > 0 {
		m["user_id"] = uid
	} else {
		err = errors.New(conf.IniConf.String("userNotExist"))
	}
	return m, err

}
