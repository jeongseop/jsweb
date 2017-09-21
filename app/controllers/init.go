package controllers

import (
	"github.com/revel/revel"
	"time"
	"fmt"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ContinuousCharter(s, c string) (int, int) {
	start := strings.Index(s, c)
	if start == -1 {
		return -1, -1
	}

	var end int
	for t := range s[start:] {
		if s[start] != s[t] {
			end = t-1
		}
	}
	return start, end
}

func getRightString(s string, c int) string {
	l := len(s)
	if l < c {
		return s
	}
	if c <= 0 {
		return s
	}
	return s[l-c:]
}

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

	revel.TemplateFuncs["substr"] = func(str string, length int) string {
		if len(str) < length {
			return str
		}
		return str[:length]
	}

	revel.TemplateFuncs["getStringFromUnix"] = func(unix int64) string {
		y, m, d := time.Unix(unix, 0).Date()
		return fmt.Sprintf("%04d%02d%02d",y,m,d)
	}

	revel.TemplateFuncs["replace"] = func(s, old, new string) string {
		return strings.Replace(s, old, new,-1)
	}

	revel.TemplateFuncs["split"] = func(s, f string) []string {
		return strings.Split(s, f)
	}
}
