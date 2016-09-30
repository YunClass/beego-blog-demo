package main

import (
	"github.com/astaxie/beego/orm"
	"myapp/models"

	"github.com/astaxie/beego"
	_ "myapp/routers"
)

func init() {
	models.RegisterDB()
}

func main() {

	orm.Debug = true
	//是否打印相关信息
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
