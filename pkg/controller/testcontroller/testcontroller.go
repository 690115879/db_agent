package controller

import (
	"db_agent/pkg/global"

	"github.com/kataras/iris/v12/mvc"
)

func init() {
	testParty := global.Webapp.Party("/test")
	m := mvc.New(testParty)
	m.Handle(new(TestController))
}

type TestController struct {
}

// GET: http://localhost:12580/test
func (c *TestController) Get() string {
	return "helloWorld!"
}
