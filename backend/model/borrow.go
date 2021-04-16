package model

import (
	"time"
)

type Borrow struct {
	ID         uint   `gorm:"not null;autoIncrement;primaryKey"`
	BookID     string `gorm:"not null"`
	Book       Book
	CardID     string `gorm:"not null"`
	Card       Card
	BorrowDate time.Time `gorm:"not null"`
	ReturnDate time.Time
}
