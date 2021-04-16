package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Book struct {
	ID        uint            `gorm:"not null;autoIncrement;primaryKey"`
	BNo       string          `gorm:"size:15;not null"`
	Category  string          `gorm:"size:15;not null"`
	Title     string          `gorm:"size:31;not null"`
	Press     string          `gorm:"size:31;not null"`
	Year      int             `gorm:"not null"`
	Author    string          `gorm:"size:15;not null"`
	Price     decimal.Decimal `gorm:"type:decimal(7,2);not null"`
	Total     uint            `gorm:"not null"`
	Stock     uint            `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
