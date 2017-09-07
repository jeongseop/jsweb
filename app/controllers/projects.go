package controllers

import (
	"github.com/jeongseop/jsweb/app/models"
	"github.com/jeongseop/jsweb/app/routes"
	"github.com/revel/revel"
	"strings"
	"fmt"
	"errors"
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
	res := make(map[string] interface{})
	memb := c.connected()
	if memb == nil {
		res["message"] = "not authorize"
		c.Response.Status = 401
		return c.RenderJSON(res)
	}

	project.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		res["message"] = c.Validation.Errors[0].String()
		c.Response.Status = 500
		return c.RenderJSON(res)
	}

	if err := c.Txn.Insert(&project); err != nil {
		res["message"] = "Insert Error!!"+ err.Error()
		c.Response.Status = 500
		return c.RenderJSON(res)
	}
	res["message"] = "Add New Project Success!!"

	return c.RenderJSON(res)
}
func (c Projects) Update(id int, project models.Project) revel.Result {
	res := make(map[string] interface{})
	memb := c.connected()
	if memb == nil {
		res["message"] = "not authorize"
		c.Response.Status = 401
		return c.RenderJSON(res)
	}
	project.ProjectId = id

	project.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		res["message"] = c.Validation.Errors[0].String()
		c.Response.Status = 500
		return c.RenderJSON(res)
	}

	cnt, err := c.Txn.Update(&project)
	if err != nil {
		res["message"] = "Update Failed.. [" + err.Error() + "]"
		c.Response.Status = 500
		return c.RenderJSON(res)
	}
	res["message"] = fmt.Sprintf("Update Success! [%d] rows", cnt)
	return c.RenderJSON(res)
}
func (c Projects) Delete(id int, project models.Project) revel.Result {
	res := make(map[string] interface{})
	memb := c.connected()
	if memb == nil {
		res["message"] = "not authorize"
		c.Response.Status = 401
		return c.RenderJSON(res)
	}
	project.ProjectId = id

	cnt, err := c.Txn.Delete(&project)
	if err != nil {
		res["message"] = "Delete Failed.. [" + err.Error() + "]"
		c.Response.Status = 500
		return c.RenderJSON(res)
	}
	res["message"] = fmt.Sprintf("Delete Success! [%d] rows", cnt)
	return c.RenderJSON(res)
}
func (c Projects) AddForm(id int) revel.Result {
	memb := c.connected()
	if memb == nil {
		c.Response.Status = 401
		return c.RenderError(errors.New("Unauthenticated"))
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