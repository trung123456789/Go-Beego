package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Employee struct {
	Id       int `orm:"pk;auto"`
	Name     string
	Age      int
	Address  string
	Position string
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime);null"`
}

func init() {
	orm.RegisterModel(new(Employee))
}

func (e *Employee) TableName() string {
	return "employee"
}