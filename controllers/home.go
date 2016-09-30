package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "home.html"
	this.Data["IsHome"] = true

	this.Data["IsLogin"] = checkLogin(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopic(true)

	if err != nil {
		beego.Error(err)
	}
}
