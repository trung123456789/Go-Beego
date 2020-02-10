package connection

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

// CreateConnection function
func CreateConnection() (*sql.DB, error) {
	var err error
	host := beego.AppConfig.String("dbHost")
	port := beego.AppConfig.String("dbPort")
	user := beego.AppConfig.String("dbUser")
	pass := beego.AppConfig.String("dbPass")
	name := beego.AppConfig.String("dbName")
	driver := beego.AppConfig.String("driver")
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)
	db, err := sql.Open(driver, sqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
