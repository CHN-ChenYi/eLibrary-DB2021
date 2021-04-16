package model

type Card struct {
	CardID     string `gorm:"primaryKey;size:15;not null"`
	Department string `gorm:"size:63;not null"`
	Type       string `gorm:"size:1;not null;check:type in ('S','T')"`
}
