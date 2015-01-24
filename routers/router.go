package routers

import (
	"github.com/pravj/hackman/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/logout", &controllers.MainController{}, "get:Logout")
    beego.Router("/callback", &controllers.OauthController{}, "get:ParseCode")
    beego.Router("/admin", &controllers.AdminController{})
    beego.Router("/admin/hackathon/?:hackathon", &controllers.AdminController{}, "get:AdminEvent")
    beego.Router("/organize", &controllers.OrganizeController{}, "post:Organize")

    beego.Router("/announce/?:category", &controllers.AnnounceController{}, "post:Announce")
    beego.Router("/public", &controllers.PublicController{})

    beego.Router("/team", &controllers.TeamController{})
    beego.Router("/hackathon", &controllers.HackathonController{})

}
