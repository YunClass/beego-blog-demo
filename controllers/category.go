package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	op := this.GetString("op")
	var err error
	if op == "del" {

		id := this.GetString("id")
		cid, _ := strconv.ParseInt(id, 10, 64)
		err = models.DelCategory(cid)
	}

	this.TplName = "category.html"
	this.Data["IsCategory"] = true
	this.Data["Categories"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
	this.Data["IsLogin"] = checkLogin(this.Ctx)

}

func (this *CategoryController) Post() {

	op := this.GetString("op")
	if op == "add" {

		name := this.GetString("name")
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
	}

	this.Redirect("/category", 301)
	return
}
