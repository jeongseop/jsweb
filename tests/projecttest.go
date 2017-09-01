package tests

import (
	"github.com/revel/revel/testing"
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
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *ProjectTest) After() {
	println("Tear down")
}
