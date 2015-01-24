package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
        v := this.GetSession("hackman")
        if v == nil {
          this.Redirect("/", 302)
          return
        } else {
          w, _ := v.(map[string]string)
          this.Data["Name"] = w["name"]
          this.Data["Avatar"] = w["avatar"]

          if w["profile"] == "admin" {
            this.TplNames = "admin.tpl"
          } else if w["profile"] == "user" {
            this.Redirect("/", 302)
            return
          }
        }
}
