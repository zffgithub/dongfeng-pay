package main

import (
	_ "agent/models"
	_ "agent/routers"
	"agent/utils"
	"log"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	// 初始化Session
	// sys.InitSession()
	utils.InitLogs()
	// 如果是开发模式，则显示命令信息
	s, _ := beego.AppConfig.String("runmode")
	isDev := !(s != "dev")
	if isDev {
		orm.Debug = isDev
	}
}

func main() {
	// web.BConfig.WebConfig.Session.SessionOn = true
	log.Println(123)
	log.Println(beego.BConfig.WebConfig.Session.SessionOn)
	beego.Run()
}
