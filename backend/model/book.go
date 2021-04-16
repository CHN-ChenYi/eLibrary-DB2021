package model

type Book struct {
	BookID   string  `gorm:"primaryKey;size:15;not null" json:"book_id" validate:"required"`
	Category string  `gorm:"size:15;not null" json:"category" validate:"required"`
	Title    string  `gorm:"size:31;not null" json:"title" validate:"required"`
	Press    string  `gorm:"size:31;not null" json:"press" validate:"required"`
	Year     int     `gorm:"not null" json:"year" validate:"required"`
	Author   string  `gorm:"size:15;not null" json:"author" validate:"required"`
	Price    float32 `gorm:"type:decimal(7,2);not null" json:"price" validate:"required"`
	Total    uint    `gorm:"not null" json:"total" validate:"required"`
	Stock    uint    `gorm:"not null;check:stock<=total" validate:"required"`
}

func CreateBook(book *Book) error {
	result := gormDb.Exec(`INSERT INTO books(book_id, category, title, press, year, author, price, total, stock) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		book.BookID, book.Category, book.Title, book.Press, book.Year, book.Author, book.Price, book.Total, book.Stock)
	return result.Error
}

func ModifyBook(book *Book) error {
	result := gormDb.Exec(`UPDATE books SET category = ?, title = ?, press = ?, year = ?, author = ?, price = ?, total = ?, stock = ? WHERE book_id = ?`,
		book.Category, book.Title, book.Press, book.Year, book.Author, book.Price, book.Total, book.Stock, book.BookID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

func queryBook(sql string, values ...interface{}) ([]Book, error) {
	ret := make([]Book, 0)
	var book Book
	rows, err := gormDb.Raw(sql, values...).Rows()
	if err != nil {
		return ret, err
	}
	defer rows.Close()
	for rows.Next() {
		gormDb.ScanRows(rows, &book)
		ret = append(ret, book)
	}
	return ret, err
}

func QueryBookAll() ([]Book, error) {
	return queryBook("SELECT * FROM books")
}

func QueryBookByCategory(category string) ([]Book, error) {
	return queryBook("SELECT * FROM books WHERE category = ?", category)
}

func QueryBookByTitle(title string) ([]Book, error) {
	return queryBook("SELECT * FROM books WHERE title = ?", title)
}

func QueryBookByPress(press string) ([]Book, error) {
	return queryBook("SELECT * FROM books WHERE press = ?", press)
}

func QueryBookByYear(l, r int) ([]Book, error) {
	return queryBook("SELECT * FROM books WHERE ? <= year and year <= ?", l, r)
}

func QueryBookByAuthor(author string) ([]Book, error) {
	return queryBook("SELECT * FROM books WHERE author = ?", author)
}

func QueryBookByPrice(l, r float32) ([]Book, error) {
	return queryBook("SELECT * FROM books WHERE ? <= price and price <= ?", l, r)
}
