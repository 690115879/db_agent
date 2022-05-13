package main

import (
	_ "db_agent/pkg/orm"
	_ "db_agent/pkg/webframework"
	_ "db_agent/pkg/controller"
	"db_agent/pkg/global"
	"fmt"
)

func main() {
	fmt.Println("db_agent start")
	global.Webapp.Listen(":12580")
}
