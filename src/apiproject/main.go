package main

import (
	_ "apiproject/routers"
	"github.com/astaxie/beego/session"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	sessionConf := &session.ManagerConfig{
		CookieName:     "begoosessionID",
		Gclifetime:     2,
		Maxlifetime:    10,
		CookieLifeTime: 10,
	}

	globalSessions, _ := session.NewManager("memory", sessionConf)
	go globalSessions.GC()
	beego.Run()
}
