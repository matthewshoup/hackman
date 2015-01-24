package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/pravj/hackman/models"
	"io/ioutil"
	"net/http"
)

const (
	TOKEN_ENDPOINT string = "https://github.com/login/oauth/access_token"
        USER_ENDPOINT string = "https://api.github.com/user"
)

type OauthController struct {
	beego.Controller
}

type payload struct {
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type response struct {
	AccessToken string `json:"access_token"`
}

type credential struct {
        Name string `json:"name"`
        UserName string `json:"login"`
        Email string `json:"email"`
        Avatar string `json:"avatar_url"`
}

func (this *OauthController) ParseCode() {
	token := AccessToken(this.GetString("code"), beego.AppConfig.String("client_id"), beego.AppConfig.String("client_secret"))
        name, username, email, avatar := Credentials(token)

	user := models.User{Token: token, Name: name, UserName: username, Email: email, Avatar: avatar, Admin: "yes"}

        ss := make(map[string]string)
        ss["email"] = email
        ss["name"] = name
        ss["username"] = username
        ss["avatar"] = avatar

        isAdmin := models.CreateUser(&user)

        if isAdmin {
          ss["profile"] = "admin"
          this.SetSession("hackman", ss)

          beego.Info("moving to admin")
          this.Redirect("/admin", 302)
        } else {
          ss["profile"] = "user"
          this.SetSession("hackman", ss)

          beego.Info("moving to user")
          this.Redirect("/", 302)
        }
        return
}

func Credentials(AccessToken string) (string, string, string, string) {
        req, _ := http.NewRequest("GET", USER_ENDPOINT, nil)

        AuthHeader := "token " + AccessToken
        req.Header.Set("Authorization", AuthHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

        client := &http.Client{}
        res, _ := client.Do(req)

        defer res.Body.Close()
        body, _ := ioutil.ReadAll(res.Body)

        var cred credential
        json.Unmarshal(body, &cred)

        return cred.Name, cred.UserName, cred.Email, cred.Avatar
}

func AccessToken(Code, ClientId, ClientSecret string) string {
	payloadJson, _ := json.Marshal(payload{Code, ClientId, ClientSecret})
	payloadReader := bytes.NewReader(payloadJson)

	req, _ := http.NewRequest("POST", TOKEN_ENDPOINT, payloadReader)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var resp response
	json.Unmarshal(body, &resp)

	return resp.AccessToken
}
