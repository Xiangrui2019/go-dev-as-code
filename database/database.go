package database

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xiangrui2019/go-dev-as-code/models"
	"github.com/xormplus/xorm"
)

var XEngine *xorm.Engine

func init() {
	var err error

	XEngine, err = xorm.NewEngine("mysql", os.Getenv("MYSQL_DSN"))

	if err != nil {
		log.Fatal(err)
	}

	MigrationDatabase()
}

func MigrationDatabase() {
	XEngine.Sync2(new(models.User))
}
