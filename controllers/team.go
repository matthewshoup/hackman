package controllers

import (
	"github.com/astaxie/beego"
)

// oprations for Team
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

}

func (c *TeamController) GetOne() {

}

func (c *TeamController) GetAll() {
	
}

func (c *TeamController) Put() {

}

func (c *TeamController) Delete() {

}
