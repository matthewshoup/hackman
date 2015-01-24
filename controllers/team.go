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

        /*Create team first, then add boys.*/
        
        /*Get UserId of CUrrent User*/
        Uid1, _ := models.GetUserByUsername(username)
        team, err := models.GetTeamByName(name)

        /*Check wheather team exists*/
        if err != nil {
            /*Update the entries*/
            Uid, _ := models.GetUserByEmail(email) //Add error handling for this
            if(team.UserId2 == -1){
                team.UserId2 = Uid.Id
            }else if(team.UserId3 == -1){
                team.UserId3 = Uid.Id
            }else if(team.UserId4 == -1){
                team.UserId4 = Uid.Id
            }

            err1 := models.UpdateTeamById(team)
            if err1 == nil {
                //SendMail(email, hackathonName, teamName)
            }
        } else {
            /*Create New team */
            team := models.Team{
                Name: name, 
                UserId1: Uid1.Id, 
                UserId2: -1, 
                UserId3: -1, 
                UserId4: -1, 
                CreatorId: Uid1.Id,  
                AccByU1:true, 
                AccByU2:false, 
                AccByU3: false, 
                AccByU4: false, 
                HackathonId: hackathonId, 
                CreatedAt: time.Now().Local(),
            }
            models.AddTeam(&team)
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
