package main

import (
	"learn/go-xorm/pkg/controllers"
	"learn/go-xorm/pkg/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	// create sql database
	if err := dao.ConnectDatabase(); nil != err {
		panic(err)
	}
	defer dao.DB.Close() // program exit close db connection

	r := gin.Default()

	r.POST("/book/", controllers.CreateBook)

	r.Run() // listen and serve on http://localhost:8080/
}
