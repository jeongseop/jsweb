package controllers

import (
	"github.com/jeongseop/jsweb/app/models"
	"github.com/jeongseop/jsweb/app/routes"
	"github.com/revel/revel"
	"log"
	"time"
	"strings"
	"fmt"
)

type Projects struct {
	App
}

func (c Projects) Project(id int) revel.Result {
	memb := c.connected()
	if memb != nil {
		c.ViewArgs["user"] = memb
	}

	if id != -1 {
		var project *models.Project
		if err := c.Txn.SelectOne(&project, `select * from project where id = ?`, id); err != nil {
			c.Flash.Error("올바르지 않은 접근입니다.")
			return c.Redirect(routes.App.Index())
		}

		project.CommentList = strings.Split(project.ProjectComment, fmt.Sprintf("%c%c%c%c",13,10,13,10))
		c.ViewArgs["project"] = project
	}

	if id == -1 && memb == nil {
		c.Flash.Error("Bad Request")
		return c.Redirect(routes.App.Index())
	}

	return c.Render()
}

func (c Projects) Add(project models.Project) revel.Result {
	memb := c.connected()
	if memb == nil {
		return c.NotFound("")
	}

	st, err := time.Parse("20060102",project.StartDate)
	log.Println(project.StartDate, st)
	if err != nil {
		c.Flash.Error("StartDate Parsing failed!!", err)
		return c.Redirect(routes.Projects.AddForm(0))
	}
	project.StartDateTime = st.UnixNano()

	ed, err := time.Parse("20060102",project.EndDate)
	log.Println(project.EndDate, ed)
	if err != nil {
		c.Flash.Error("EndDate Parsing failed!!", err)
		return c.Redirect(routes.Projects.AddForm(0))
	}
	project.EndDateTime = ed.UnixNano()

	log.Printf("ttt[%s], [%s]\n",time.Unix(0,project.StartDateTime).String(),time.Unix(0,project.EndDateTime).String())
	project.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Index())
	}

	if err := c.Txn.Insert(&project); err != nil {
		c.Flash.Error("Insert Error!!", err)
		return c.Redirect(routes.Projects.AddForm(0))
	}
	c.Flash.Success("Add New Project Success!!")

	return c.Redirect(routes.App.Index())
}
func (c Projects) Update(id int) revel.Result {
	return c.Render()
}
func (c Projects) Delete(id int) revel.Result {
	return c.Render()
}
func (c Projects) AddForm(id int) revel.Result {
	memb := c.connected()
	if memb == nil {
		return c.NotFound("")
	}

	if id > 0 {
		var project models.Project
		if err := c.Txn.SelectOne(&project, `select * from project where id = ?`,id); err != nil {
			c.Flash.Error(err.Error())
			return c.Redirect(routes.App.Index())
		}
		c.ViewArgs["project"] = project
	}
	return c.Render()
}
func (c Projects) EditForm(id int) revel.Result {
	return c.Render()
}