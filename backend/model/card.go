package model

import (
	"gorm.io/gorm"
)

type Card struct {
	CardID     string `gorm:"primaryKey;size:15;not null" json:"card_id" validate:"required"`
	Department string `gorm:"size:63;not null" json:"department" validate:"required"`
	Type       string `gorm:"size:1;not null;check:type in ('S','T')" json:"type" validate:"required,oneof=S T"`
	Deleted    bool   `gorm:"not null"`
}

func CreateCard(card *Card) error {
	result := gormDb.Exec(`INSERT INTO cards(card_id, department, type, deleted) VALUES (?, ?, ?, ?)`, card.CardID, card.Department, card.Type, card.Deleted)
	return result.Error
}

func ModifyCard(card *Card) error {
	result := gormDb.Exec(`UPDATE cards SET department = ?, type = ? WHERE card_id = ? AND deleted = 0`, card.Department, card.Type, card.CardID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

func DeleteCard(cardID string) error {
	result := gormDb.Exec(`UPDATE cards SET deleted = true WHERE card_id = ?`, cardID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

func ValidateCard(cardID string) (bool, error) {
	var card Card
	rows, err := gormDb.Raw("SELECT * FROM cards WHERE card_id = ?", cardID).Rows()
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		gormDb.ScanRows(rows, &card)
		return !card.Deleted, nil
	}
	return false, gorm.ErrRecordNotFound
}
