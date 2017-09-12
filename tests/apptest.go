package tests

import (
	"github.com/revel/revel/testing"
	"net/url"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestLoginForm() {
	t.Get("/login")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
func (t *AppTest) TestLogin() {
	t.PostForm("/login", url.Values{"id":{"jeongseop"}, "password":{"demo"}})

	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
func (t *AppTest) TestLogout() {
	t.Get("/logout")

	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}
