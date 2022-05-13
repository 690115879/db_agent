package orm

import (
	"db_agent/pkg/global"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	var err error
	global.DataDb, err = gorm.Open(sqlite.Open("./data/data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect db")
	}
}
