package routes

import (
	"learn/go-xorm/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(r *gin.Engine) {
	r.POST("/book/", controllers.CreateBook)
}
