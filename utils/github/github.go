package main

import "reflect"
import "fmt"
import "bytes"
import "encoding/json"
import "github.com/pravj/hackman/utils/request"

const (
  TEAM_CREATION_ENDPOINT string = "https://api.github.com/orgs/"
)

type User struct {
  UserName string
}

type Team struct {
  Id int
  Name string
  Users []User
}

type TeamParameter struct {
  Name string `json:"name"`
  Permission string `json:"permission"`
}

type TeamResponse struct {
  Id int `json:"id"`
}

type Organization struct {
  Name string
  Teams []Team
}

func CreateTeams(org *Organization, accessToken string) {
  for _, team := range org.Teams {
    payloadJson, _ := json.Marshal(TeamParameter{Name: team.Name, Permission: "push"})
    payloadReader := bytes.NewReader(payloadJson)

    body := request.Post(TEAM_CREATION_ENDPOINT + org.Name + "/teams", accessToken, payloadReader)
    fmt.Println(string(body))

    var tr TeamResponse
    json.Unmarshal(body, &tr)

    fmt.Println(tr.Id)
  }
}

func main() {
  user1 := User{UserName: "nkman"}
  user2 := User{UserName: "iMshyam"}

  fmt.Println(reflect.TypeOf(user1))

  team := Team{Id: 0, Name: "test", Users: []User{user1, user2}}
  fmt.Println(reflect.TypeOf(team))
  org := Organization{Name: "mockers", Teams: []Team{team}}

  CreateTeams(&org, "xxxxxxxxxxxx")
}
