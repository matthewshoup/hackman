package controllers

import (
    "github.com/astaxie/beego"
    "github.com/pravj/hackman/models"
    "time"
    "strconv"
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

func (this *TeamController) Post() {
    v := this.GetSession("hackman")
    if v == nil {
        this.TplNames = "index.tpl"
    }else {
        w, _ := v.(map[string]string)

        username := w["username"]
        name := this.Input().Get("name")
        email := this.Input().Get("email")
        hackathonId, _ := strconv.Atoi(this.Input().Get("hackathonId"))

        if email != "" {
            Uid1, _ := models.GetUserByUsername(username)
            Uid, err := models.GetUserByEmail(email)
            /*Error handling here*/

            team, err := models.GetTeamByName(name)
            if err != nil {
                //Update the entries/
                beego.Info(team.Id)
            } else{
                //Create the team entry in team table.
                team := models.Team{Name: name, UserId1: Uid1.Id, UserId2: Uid.Id, UserId3: 0, UserId4: 0, CreatorId: Uid1.Id,  AccByU1:true, AccByU2:false, AccByU3: false, AccByU4: false, HackathonId: hackathonId, CreatedAt: time.Now().Local()}
                models.AddTeam(&team)
            }
        } else{
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
