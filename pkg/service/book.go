package service

import (
	"fmt"
	"learn/go-xorm/pkg/dao"
	"learn/go-xorm/pkg/models"
)

// CreateBook inserts a new book to d
func CreateBook(book *models.Book) (err error) {
	affected, err := dao.DB.Insert(book)
	if nil != err {
		return
	}
	fmt.Println(affected)
	return nil
}
