package tests

import (
	"github.com/revel/revel/testing"
	"net/url"
)

type ProjectTest struct {
	testing.TestSuite
}

func (t *ProjectTest) Before() {
	println("Set up")
}

func (t *ProjectTest) TestShowProject() {
	t.Get("/projects/1")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *ProjectTest) TestAddProjectForm() {
	t.Get("/projects/new")
	t.AssertNotFound()
}

func (t *ProjectTest) TestAddProject() {
	t.PostForm("/projects",url.Values{"project.ProjectName":{"test"},"project.ProjectComment":{"test11"},"project.CompanyName":{"asdf"},"project.StartDate":{"20170904"},"project.EndDate":{"20170910"},"project.LaunchUrl":{"test111"}})
	t.AssertNotFound()
}

func (t *ProjectTest) After() {
	println("Tear down")
}
