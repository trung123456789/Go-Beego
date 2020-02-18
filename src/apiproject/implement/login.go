package implement

import (
	"apiproject/models"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

func Login(id, password string) (bool, error) {

	o := orm.NewOrm()
	o.Using("default")

	// Get a QuerySeter object. User is table name
	qs := o.QueryTable("userinfo")
	userInfo := models.UserInfo{}
	errSelect := qs.Filter("UserId", id).One(&userInfo)
	if errSelect != nil {
		return false, errSelect
	}
	if !reflect.DeepEqual(userInfo, models.UserInfo{}) {
		checkPass := CheckPasswordHash(password, userInfo.Password)
		if checkPass {
			return  true, nil
		}
	}
	return false, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
