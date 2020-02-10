package conf

import (
	"github.com/astaxie/beego/config"
	"log"
)

var (
	IniConf config.Configer
	err error
)

func init()  {
	IniConf, err = config.NewConfig("ini", "./conf/msg.conf")
	if IniConf != nil {
		log.Println(err)
	}
}
