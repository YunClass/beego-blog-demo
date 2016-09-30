package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
	"strconv"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {

	this.TplName = "topic.html"

	var err error
	this.Data["Topics"], err = models.GetAllTopic(false)

	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Post() {

	if !checkLogin(this.Ctx) {

		this.Redirect("/login", 302)
		return
	}

	title := this.GetString("title")
	content := this.GetString("content")

	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {

	idStr := this.Ctx.Input.Params()["0"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	this.TplName = "topic_view.html"
	this.Data["Topic"], err = models.GetTopic(id)
}

func (this *TopicController) Modify() {

	this.TplName = "topic_modify.html"
	idInt, err := strconv.ParseInt(this.Ctx.Input.Params()["0"], 10, 64)

	if err != nil {
		beego.Error(err)
	}

	var topic *models.Topic
	topic, err = models.GetTopic(idInt)
	if err != nil {
		beego.Error(err)
	}

	this.Data["Topic"] = topic
}
