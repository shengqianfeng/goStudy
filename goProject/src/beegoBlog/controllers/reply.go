package controllers

import (
	"beegoBlog/models"

	"github.com/astaxie/beego"
	log "github.com/keepeye/log4go"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")
	err := models.AddReply(tid,
		this.Input().Get("nickname"), this.Input().Get("content"))
	if err != nil {
		log.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	if !checkAccount(this.Ctx) {
		return
	}
	tid := this.Input().Get("tid")
	err := models.DeleteReply(this.Input().Get("rid"))
	if err != nil {
		log.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}
