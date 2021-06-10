package controllers

import (
	"beegoBlog/models"
	"path"
	"strings"

	"github.com/astaxie/beego"
	log "github.com/keepeye/log4go"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		log.Error(err)
	}
	this.Data["Topics"] = topics
}

//所有有关topic的post方法都会自动路由进入这个方法
func (this *TopicController) Post() {
	//未登录则跳转到登陆页面302
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	label := this.Input().Get("label")
	content := this.Input().Get("content")
	category := this.Input().Get("category")

	// 获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		//不需要return，附件失败还有其他信息
		log.Error(err)
	}

	var attachment string
	if fh != nil {
		// 保存附件
		attachment = fh.Filename
		log.Info(attachment)

		//第二个参数为文件保存路径
		//SaveToFile saves uploaded file to new path. it only operates the first one of mutil-upload form file field.
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			log.Error(err)
		}
	}

	//根据id判断是修改还是添加
	if len(tid) == 0 {
		err = models.AddTopic(title, category, label, content, attachment)
	} else {
		err = models.ModifyTopic(tid, title, category, label, content, attachment)
	}

	if err != nil {
		log.Error(err)
	}
	//跳转到topic.html列表
	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "topic_add.html"
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		log.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		log.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	tid := this.Ctx.Input.Param("0")
	topic, err := models.GetTopic(tid)
	if err != nil {
		log.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Lables"] = strings.Split(topic.Lables, " ")
	replies, err := models.GetAllReplies(tid)
	if err != nil {
		log.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}
