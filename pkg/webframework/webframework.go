package webframework

import (
	"db_agent/pkg/global"

	"github.com/kataras/iris/v12"
)

func init() {
	global.Webapp = iris.New()
}
