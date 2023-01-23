package controllers

import (
	"fmt"
	"learn/go-xorm/pkg/models"
	"learn/go-xorm/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBook creates a new book
func CreateBook(c *gin.Context) {
	// parse data to json from request
	var book models.Book
	c.BindJSON(&book)
	fmt.Println(book)
	// save to database
	if err := service.CreateBook(&book); nil != err {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": book,
		})
	}
}
