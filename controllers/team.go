package controllers

import (
    "github.com/astaxie/beego"
)

type TeamController struct {
    beego.Controller
}

func (c *TeamController) URLMapping() {
    c.Mapping("Post", c.Post)
    c.Mapping("GetOne", c.GetOne)
    c.Mapping("GetAll", c.GetAll)
    c.Mapping("Put", c.Put)
    c.Mapping("Delete", c.Delete)
}

func (c *TeamController) Post() {
    v := this.GetSession("hackman")
    if v == nil {
        this.TplNames = "index.tpl"
    } 
    else {
        w, _ := v.(map[string]string)
        ss := make(map[string]string)

        ss["username"] = w["username"]
        ss["name"] = c.Input().Get("name")
        ss["email1"] = c.Input().Get("email")

        if ss["email"] != nil {
            Uid, err := models.GetUserByEmail(ss["email"])
            /*Error handling here*/

            team, err := models.GetTeamByName(ss["name"])
            if err != nil{
                //Update the entries/
            }
            else{
                //Create the team entry in team table.
            }
        }
        else{
            beego.Error("Email is missing !!")
        }
    }
}

func (c *TeamController) GetOne() {

}

func (c *TeamController) GetAll() {

}

func (c *TeamController) Put() {

}

func (c *TeamController) Delete() {

}
