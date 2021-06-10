package controllers

import (
	"beegoBlog/models"

	"github.com/astaxie/beego"
	log "github.com/keepeye/log4go"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(this.Input().Get("cate"), this.Input().Get("label"), true)
	if err != nil {
		log.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		log.Error(err)
	}
	this.Data["Categories"] = categories
}
