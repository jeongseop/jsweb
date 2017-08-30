package controllers

import (
	"github.com/jeongseop/jsweb/app/models"
	"github.com/jeongseop/jsweb/app/routes"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type App struct {
	GorpController
}

func (c App) connected() *models.Member {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.Member)
	}
	if id, ok := c.Session["user"]; ok {
		return c.getUser(id)
	}
	return nil
}

func (c App) getProjectList() []models.Project {
	//var pList []models.Project
	//_, err := c.Txn.Select(&pList, `select * from project`)
	//if err != nil {
	//	panic(err)
	//}
	//return pList

	t := make([]models.Project, 2)
	t[0] = models.Project{0,"test1","test111111","comp1","web","20170830","20170830",time.Now(),time.Now()}
	t[1] = models.Project{1,"test2","test222222","comp2","server","20170830","20170830",time.Now(),time.Now()}
	return t
}

func (c App) Index() revel.Result {
	memb := c.connected()
	if memb != nil {
		c.ViewArgs["user"] = memb
	}

	//Project List
	var projList []models.Project
	projList = c.getProjectList()
	c.ViewArgs["portfolio_list"] = projList

	return c.Render()
}

func (c App) Blog() revel.Result {
	return c.Render()
}

func (c App) LoginForm() revel.Result {
	memb := c.connected()
	if memb != nil {
		return c.Redirect(routes.App.Index())
	}
	return c.Render()
}

func (c App) getUser(id string) *models.Member {
	var m *models.Member
	//err := c.Txn.SelectOne(&m, `select * from member where id = ?`, id)
	//if err != nil {
	//	panic(err)
	//}
	bcryptPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("demo"), bcrypt.DefaultCost)
	m = &models.Member{"jeongseop", "demo", "asdf@asdf.com", bcryptPassword}
	return m
}

func (c App) Login(id, password string) revel.Result {
	if c.connected() != nil {
		c.Redirect(routes.App.Index())
	}

	member := c.getUser(id)
	if member != nil {
		err := bcrypt.CompareHashAndPassword(member.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = id
			c.Session.SetNoExpiration()
			c.Flash.Success("Welcome, " + id)
			return c.Redirect(routes.App.Index())
		}
	}

	c.Flash.Out["id"] = id
	c.Flash.Error("Login failed")
	return c.Redirect(routes.App.LoginForm())
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}

func (c App) Page() revel.Result {
	return c.Render()
}
func (c App) Icons() revel.Result {
	return c.RenderTemplate("App/page-icons.html")
}
func (c App) Elements() revel.Result {
	return c.RenderTemplate("App/page-elements.html")
}
func (c App) Typography() revel.Result {
	return c.RenderTemplate("App/page-typography.html")
}
func (c App) Sidebar() revel.Result {
	return c.RenderTemplate("App/page-sidebar.html")
}
func (c App) Contact() revel.Result {
	return c.Render()
}
func (c App) Project() revel.Result {
	return c.Render()
}
func (c App) Portfolio() revel.Result {
	return c.Render()
}
func (c App) Single() revel.Result {
	return c.Render()
}
