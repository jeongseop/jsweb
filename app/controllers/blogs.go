package controllers

import (
	"github.com/revel/revel"
	"errors"
)

type Blogs struct {
	App
}

const (
	DEFAULT_IMAGE_PATH="/public/upload/default/"
)

func (c Blogs) List() revel.Result {
	memb := c.connected()
	if memb != nil {
		c.ViewArgs["user"] = memb
	}

	return c.Render()
}

func (c Blogs) View() revel.Result {
	memb := c.connected()
	if memb != nil {
		c.ViewArgs["user"] = memb
	}
	return c.Render()
}

func (c Blogs) Add() revel.Result {
	c.Response.Status = 501
/*
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {

	}

	mpartfile, mpartHeader, err := c.Request.FormFile("upload")
	if err != nil {

	}
	defer mpartfile.Close()
	imagePath := revel.Config.StringDefault("upload.image.path",DEFAULT_IMAGE_PATH)
*/
	return c.Render()
}
func (c Blogs) Update() revel.Result {
	c.Response.Status = 501
	return c.Render()
}
func (c Blogs) Delete() revel.Result {
	c.Response.Status = 501
	return c.Render()
}
func (c Blogs) AddForm() revel.Result {
	memb := c.connected()
	if memb == nil {
		c.Response.Status = 401
		return c.RenderError(errors.New("Unauthenticated"))
	}
	//c.Response.Status = 501
	return c.RenderTemplate("Blogs/write.html")
}