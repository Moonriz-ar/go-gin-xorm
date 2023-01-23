package dao

import (
	"fmt"
	"learn/go-xorm/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var DB *xorm.Engine

var (
	userName string = "root"
	password string = "password"
	dbName   string = "db1"
	charset  string = "utf8"
)

// ConnectDatabase connects to mysql db, gets and sets table schema
func ConnectDatabase() (err error) {
	// connect to database
	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=%s", userName, password, dbName, charset)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if nil != err {
		return
	}
	fmt.Println("sql connected successfully")

	// set name mapping to GonicMapper
	engine.SetMapper(names.GonicMapper{})

	// sincronize table with struct
	if err := engine.Sync(new(models.Book)); nil != err {
		fmt.Println("error with database schema synchronize")
	}

	DB = engine
	return nil
}
