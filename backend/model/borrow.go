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

func QueryBorrowWithoutReturnDateAll(cardID string) ([]Book, error) {
	ret := make([]Book, 0)
	var record Borrow
	rows, err := gormDb.Raw("SELECT * FROM borrows WHERE card_id = ? AND return_date IS NULL", cardID).Rows()
	if err != nil {
		return ret, err
	}
	defer rows.Close()
	for rows.Next() {
		gormDb.ScanRows(rows, &record)
		book, bookErr := QueryBookByBookID(record.BookID)
		if bookErr != nil {
			return ret, err
		}
		ret = append(ret, book...)
	}
	return ret, err
}

func CreateBorrow(cardID, bookID string) error {
	result := gormDb.Exec("INSERT INTO borrows(book_id, card_id, borrow_date) VALUE(?, ?, ?)", bookID, cardID, time.Now()).Error

	if result != nil {
		return result
	}

	return nil
}

func DeleteBorrow(cardID, bookID string) error {
	result := gormDb.Exec("UPDATE borrows SET return_date = ? WHERE card_id = ? AND book_id = ? AND return_date IS NULL ORDER BY borrow_date ASC LIMIT 1", time.Now(), cardID, bookID)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}

	return nil
}
