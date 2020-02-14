package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserInfo struct {
	Id        int       `orm:"pk;auto" json:"id"`
	UserId    string    `orm:"unique" json:"user_id"`
	Name      string    `json:"user_name"`
	Age       int       `json:"age"`
	Password  string    `json:"password"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);null" json:"updated_at"`
}

func init() {
	orm.RegisterModel(new(UserInfo))
}

func (e *UserInfo) TableName() string {
	return "userinfo"
}
