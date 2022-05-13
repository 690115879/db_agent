package global

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var Webapp *iris.Application

var DataDb *gorm.DB
