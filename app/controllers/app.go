package controllers

import (
	"github.com/revel/revel"
	"myweb/app/models"
	"golang.org/x/crypto/bcrypt"
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
			return c.Render()
		}
	}

	c.Flash.Out["id"] = id
	c.Flash.Error("Login failed")
	return c.Render()
}
