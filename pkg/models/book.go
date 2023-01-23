package models

import (
	"time"
)

// Book has book title and author
type Book struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
}

func (b *Book) TableName() string {
	return "book"
}
