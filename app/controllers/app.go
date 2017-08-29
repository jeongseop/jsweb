package controllers

import (
	"github.com/jeongseop/jsweb/app/models"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"github.com/jeongseop/jsweb/app/routes"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Blog() revel.Result {
	return c.Render()
}

func (c App) LoginForm() revel.Result {
	return c.Render()
}

func (c App) getUser(id string) *models.Member {
	users, err := c.Txn.Select(models.Member{}, `select * from member where id = ?`, id)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.Member)
}

func (c App) Login(id, password string) revel.Result {
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