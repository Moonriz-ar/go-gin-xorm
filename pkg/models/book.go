package models

import (
	"time"
)

// Book has book title and author
type Book struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	Updated   time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"` // soft delete, shows delete time instead of really deleting
}

func (b *Book) TableName() string {
	return "book"
}
