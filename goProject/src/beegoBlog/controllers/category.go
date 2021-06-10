package controllers

import (
	"beegoBlog/models"

	"github.com/astaxie/beego"
	log "github.com/keepeye/log4go"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	// 检查是否有操作
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			log.Error(err)
		}

		this.Redirect("/category", 302)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DeleteCategory(id)
		if err != nil {
			log.Error(err)
		}

		this.Redirect("/category", 302)
		return
	}

	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		log.Error(err)
	}
}
