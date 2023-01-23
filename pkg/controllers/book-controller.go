package controllers

import (
	"fmt"
	"learn/go-xorm/pkg/dao"
	"learn/go-xorm/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRecentBooks returns an array of 10 recently created books
func GetRecentBooks(c *gin.Context) {
	var books []models.Book
	// query database
	if err := dao.DB.Desc("created_at").Limit(10).Find(&books); nil != err {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": books,
		})
	}
}

// GetBookByID returns the book query by id
func GetBookByID(c *gin.Context) {
	book := new(models.Book)
	// parse path param id
	id := c.Param("id")
	// query database
	affected, err := dao.DB.ID(id).Get(book)
	if nil != err {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	// could not query book with id, book could not be found
	if !affected {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": book,
		})
	}
}

// CreateBook creates a new book
// Todo: how to return to user created book with correct id (autoincremented)? currently always returns 0
func CreateBook(c *gin.Context) {
	// parse data from request to book struct, bind JSON
	var book models.Book
	c.BindJSON(&book)
	fmt.Println(book)
	// save to database
	affected, err := dao.DB.Insert(book)
	fmt.Println("CreateBook affected:", affected, book.Id)
	if nil != err {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": book,
		})
	}
}
