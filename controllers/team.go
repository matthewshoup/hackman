package controllers

import (
	"github.com/astaxie/beego"
	"github.com/pravj/hackman/models"
	"strconv"
	"time"
)

type TeamController struct {
	beego.Controller
}

type TeamConfirmController struct {
	beego.Controller
}

type messageParameter struct {
	email       string
	hackathonId int
	teamName    string
	//teamId int64
	username string
}

func (c *TeamConfirmController) URLMapping() {
	c.Mapping("Get", c.Get)
}

func (c *TeamController) URLMapping() {
	c.Mapping("Post", c.Post)
	// c.Mapping("GetOne", c.GetOne)
	// c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

func (this *TeamController) Post() {
	v := this.GetSession("hackman")
	if v == nil {
		this.TplNames = "index.tpl"
	} else {
		w, _ := v.(map[string]string)

		username := w["username"]
		name := this.Input().Get("name")
		email := this.Input().Get("email")
		hackathonId, _ := strconv.Atoi(this.Input().Get("hackathonId")) //Add error handling for this

		/*Create team first, then add boys.*/

		/*Get UserId of CUrrent User*/
		Uid1, _ := models.GetUserByUsername(username)
		team, err := models.GetTeamByName(name)

		/*Check wheather team exists*/
		if err != nil {
			/*Update the entries*/
			user, _ := models.GetUserByEmail(email) //Add error handling for this
			if team.UserId2 == -1 {
				team.UserId2 = user.Id
			} else if team.UserId3 == -1 {
				team.UserId3 = user.Id
			} else if team.UserId4 == -1 {
				team.UserId4 = user.Id
			}

			err1 := models.UpdateTeamById(team)
			if err1 == nil {
				//SendMail(email, hackathonName, teamName, teamId, username)
				var mParameter messageParameter
				mParameter.email = user.Email
				mParameter.hackathonId = hackathonId
				mParameter.teamName = team.Name
				//mParameter.teamId = team.Id
				mParameter.username = username
				SendMail(mParameter)
			}
		} else {
			/*Create New team */
			team := models.Team{
				Name:        name,
				UserId1:     Uid1.Id,
				UserId2:     -1,
				UserId3:     -1,
				UserId4:     -1,
				CreatorId:   Uid1.Id,
				AccByU1:     true,
				AccByU2:     false,
				AccByU3:     false,
				AccByU4:     false,
				HackathonId: hackathonId,
				CreatedAt:   time.Now().Local(),
			}
			models.AddTeam(&team)
		}

	}
}

func (this *TeamConfirmController) Get() {
	v := this.GetSession("hackman")
	if v == nil {
		this.TplNames = "index.tpl"
	} else {
		w, _ := v.(map[string]string)
		hackathonId, _ := strconv.Atoi(this.Input().Get("hackathonId"))
		teamName := this.Input().Get("teamName")
		email := w["email"]

		/*Add hackathon and team validation*/

		user, _ := models.GetUserByEmail(email)
		team, _ := models.GetTeamByName(teamName)

		if team.HackathonId == hackathonId {
			if team.UserId2 == user.Id {
				team.AccByU2 = true
			} else if team.UserId3 == user.Id {
				team.AccByU3 = true
			} else if team.UserId4 == user.Id {
				team.AccByU4 = true
			}

			err := models.UpdateTeamById(team)
			if err == nil {
				//SendMail(email, team.HackathonId, team.Name, team.Id, username)
				var mParameter messageParameter
				mParameter.email = user.Email
				mParameter.hackathonId = hackathonId
				mParameter.teamName = team.Name
				//mParameter.teamId = team.Id
				mParameter.username = w["username"]
				SendMail(mParameter)
			}
		} else {
			this.Data["Status"] = 0
			this.Data["message"] = "This team is not for this hackathon"
		}
		this.Data["Name"] = w["name"]
		this.Data["Avatar"] = w["avatar"]
		this.TplNames = "user.tpl"
	}
}

func SendMail(mParameter messageParameter) {
	email := mParameter.email
	hackathonId := mParameter.hackathonId
	teamName := mParameter.teamName
	//teamId := mParameter.teamId
	username := mParameter.username

	hackathon, _ := models.GetHackathonById(hackathonId)
	user, _ := models.GetUserByUsername(username)

	//message string
	message := "You have been added in team " + teamName + " By " + user.Name + " for " + hackathon.Name + ",\n <a href=http://localhost:8080/confirmteam?hackathonId=" + strconv.Itoa(hackathonId) + "&teamName=" + teamName + ">click here</a> to confirm your Team."
	beego.Info(message, email)
	/* TODO : Message Sending part */
}
