package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Company struct {
	Id               int `orm:"pk;auto"`
	Name             string
	Address          string
	Type             string
	NumberOfEmployee int
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime);null"`
}


func init() {
	orm.RegisterModel(new(Company))
}

func (e *Company) TableName() string {
	return "company"
}