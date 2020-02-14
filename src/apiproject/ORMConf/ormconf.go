package ORMConf

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"log"
)

func Init() {
	maxIdle := 30
	maxConn := 30
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// param 1:        Database alias. ORM will use it to switch database.
	// param 2:        driverName
	// param 3:        connection string
	// param 4 (optional):  set maximum idle connections
	// param 4 (optional):  set maximum connections (go >= 1.2)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=123456 dbname=Demo sslmode=disable", maxIdle, maxConn)

	orm.Debug = true

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		log.Println(err)
	}
}