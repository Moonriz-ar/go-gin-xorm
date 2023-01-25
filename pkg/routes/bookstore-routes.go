package routes

import (
	"learn/go-xorm/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterBookStoreRoutes registers to gin all the routes
var RegisterBookStoreRoutes = func(r *gin.Engine) {
	r.GET("/book", controllers.GetRecentBooks)
	r.GET("/book/:id", controllers.GetBookByID)
	r.POST("/book/", controllers.CreateBook)
	r.DELETE("/book/:id", controllers.DeleteBookByID)
}
