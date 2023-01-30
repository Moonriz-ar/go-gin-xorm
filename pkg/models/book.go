package models

import (
	"time"
)

// Book has book title and author
type Book struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" xorm:"varchar(25) not null"`
	Author    string    `json:"author" xorm:"varchar(25) not null"`
	Category  string    `json:"category" xorm:"varchar(25) not null"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	Updated   time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"` // soft delete, shows delete time instead of really deleting
}

func (b *Book) TableName() string {
	return "book"
}
