package orm

import (
	"db_agent/pkg/global"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	var err error
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "./data/data.db"
	}
	global.DataDb, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect db")
	}
}
