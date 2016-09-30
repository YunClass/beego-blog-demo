package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	/**this.Ctx.WriteString("appName:" + beego.AppConfig.String("appname") +
		"\r\nhttport:" + beego.AppConfig.String("httport") +
		"\r\nrunmode:" + beego.AppConfig.String("runmode"))

	//设置日志级别
	beego.Alert("this is alert")

	beego.Debug("this is debug")

	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(false)

	beego.Alert("this is alert2")

	beego.Debug("this is debug2")**/

	this.TplName = "index.tpl"

	this.Data["Email"] = "kongjun@maihaoche.com"

	user := &User{"张三", 10}

	this.Data["User"] = user

	//传数组
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	this.Data["Nums"] = nums
}

type User struct {
	Name string
	Age  int
}
