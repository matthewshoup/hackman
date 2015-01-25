package controllers

import (
	"github.com/astaxie/beego"
	"github.com/pravj/hackman/models"
)

type AnnounceController struct {
	beego.Controller
}

func (this *AnnounceController) Announce() {
	announcement := this.GetString("announcement")
	category := this.Ctx.Input.Param(":category")

	announce := models.Announcement{Category: category, Announcement: announcement}
	models.AddAnnouncement(&announce)

	this.Redirect("/admin", 302)
	return
}
