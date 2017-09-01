package controllers

import (
	"github.com/jeongseop/jsweb/app/models"
	"github.com/jeongseop/jsweb/app/routes"
	"github.com/revel/revel"
	"log"
	"time"
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
			panic(err)
		}
		c.ViewArgs["project"] = project
	}

	if id == -1 && memb == nil {
		c.Flash.Error("Bad Request")
		return c.Redirect(routes.App.Index())
	}

	return c.RenderTemplate("Projects/project.html")
}

func (c Projects) AddProject(project models.Project) revel.Result {
	log.Printf("ttt[%s], [%s]\n",time.Unix(0,project.StartDateTime).String(),time.Unix(0,project.EndDateTime).String())
	//c.Txn.Insert(project)

	return c.Render()
}
func (c Projects) UpdateProject(id int) revel.Result {
	return c.Render()
}
func (c Projects) DeleteProject(id int) revel.Result {
	return c.Render()
}
func (c Projects) AddProjectForm() revel.Result {
	return c.Render()
}
func (c Projects) EditProjectForm(id int) revel.Result {
	return c.Render()
}