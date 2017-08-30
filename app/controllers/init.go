package controllers

import "github.com/revel/revel"

func init() {
	//revel.OnAppStart(InitDB)
	//revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	//revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	//revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
	revel.TemplateFuncs["substr"] = func(str string, length int) string {
		if len(str) < length {
			return str
		}
		return str[:length]
	}
}
