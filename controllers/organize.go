package controllers

import (
	"github.com/astaxie/beego"
        "github.com/pravj/hackman/models"
)

type OrganizeController struct {
	beego.Controller
}

func (this *OrganizeController) Organize() {
  name := this.GetString("hackathon")
  desc := this.GetString("hackathon-description")
  org := this.GetString("hackathon-organization")

  hackathon := models.Hackathon{Name: name, Description: desc, Organization: org}
  models.CreateHackathon(&hackathon)

  this.Redirect("/admin", 302)
  return
}
