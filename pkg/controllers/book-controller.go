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
	books := []models.Book{}
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
	// if affected == false, could not query book with id, book could not be found
	if !affected {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Record with id %v not found!", id),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": book,
		})
	}
}

// UpdateBookByID updated book by id. Only fields that are included in request are updated, others remain the same
func UpdateBookByID(c *gin.Context) {
	// parse data from request to book struct, bind JSON
	book := new(models.Book)
	if err := c.BindJSON(&book); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// parse path param id
	id := c.Param("id")
	// query database
	affected, err := dao.DB.ID(id).Update(book)
	fmt.Println("affected update book by id", affected)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// if affected == 0, means there is no book with that id
	if affected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Record with id %v not found!", id)})
	}
	// if affected == 1, means book has been updated
	if affected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": book,
		})
	}
	fmt.Println(book)
}

// DeleteBookByID deletes a book by ID
func DeleteBookByID(c *gin.Context) {
	book := new(models.Book)
	// parse path param id
	id := c.Param("id")
	// query database
	affected, err := dao.DB.ID(id).Delete(book)
	if nil != err {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	fmt.Println("affected delete:", affected)
	// if affected == 0, means there is no book with that id
	if affected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Record with id %v not found!", id)})
	}
	// if affected == 1, means book with id has been deleted
	if affected == 1 {
		c.JSON(http.StatusNoContent, gin.H{
			"code": 204,
			"msg":  "success",
		})
	}
}

// CreateBook creates a new book
// Todo: how to return to user created book with correct id (autoincremented)? currently always returns 0
// should query db to get created book in db?
func CreateBook(c *gin.Context) {
	// parse data from request to book struct, bind JSON
	book := new(models.Book)
	if err := c.BindJSON(&book); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save to database
	affected, err := dao.DB.Insert(book)
	fmt.Println("CreateBook affected:", affected, book, book.ID)
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
