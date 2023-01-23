package main

import (
	"learn/go-xorm/pkg/dao"
	"learn/go-xorm/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// create sql database
	if err := dao.ConnectDatabase(); nil != err {
		panic(err)
	}
	defer dao.DB.Close() // program exit close db connection

	// instanciate gin
	r := gin.Default()

	routes.RegisterBookStoreRoutes(r)

	r.Run() // listen and serve on http://localhost:8080/
}
