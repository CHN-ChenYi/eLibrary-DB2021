package model

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ErrNoRowsAffected = errors.New("no rows affected")
var gormDb *gorm.DB

func Connect() {
	dsn := getDatabaseLoginInfo()
	logrus.Info("Connecting MySQL")
	var err error
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
	if gormDb == nil {
		logrus.Panic("DB is nil")
	}
	logrus.Info("MySQL connected")
}

func AutoMigrate() {
	if gormDb == nil {
		logrus.Panic("DB is nil")
	}
	err := gormDb.AutoMigrate(
		&Book{},
		&Card{},
		&Borrow{},
	)
	if err != nil {
		logrus.Panic(err)
	}
	// logrus.Warn("Execute ./model/borrow_trigger.sql after starting up the program for the first time.")
	addTrigger(`CREATE TRIGGER borrow_book BEFORE INSERT ON borrows
	FOR EACH ROW
	BEGIN
		IF 0 = (SELECT count(*) FROM books WHERE book_id = new.book_id) THEN
			SIGNAL SQLSTATE '45000' SET message_text = 'book_id invalid';
		END IF;
		IF 0 = (SELECT stock FROM books WHERE book_id = new.book_id) THEN
			SIGNAL SQLSTATE '45000' SET message_text = 'out of stock';
		END IF;
		UPDATE books SET stock = stock - 1 WHERE book_id = new.book_id;
	END`)
	addTrigger(`CREATE TRIGGER return_book BEFORE UPDATE ON borrows
	FOR EACH ROW
	BEGIN
		IF old.return_date IS NULL and new.return_date IS NOT NULL
			 AND old.id = new.id AND old.book_id = new.book_id
			 AND old.card_id = new.card_id AND old.borrow_date = new.borrow_date THEN
			UPDATE books SET stock = stock + 1 WHERE book_id = new.book_id;
		END IF;
	END`)
}

func getDatabaseLoginInfo() string {
	loginInfo := viper.GetStringMapString("sql")

	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?tls=skip-verify&parseTime=true&loc=Asia%%2FShanghai",
		loginInfo["user"],
		loginInfo["password"],
		loginInfo["protocol"],
		loginInfo["host"],
		loginInfo["port"],
		loginInfo["db_name"])
}

func addTrigger(trigger string) {
	err := gormDb.Exec(trigger).Error
	if err != nil && err.Error() != "Error 1359: Trigger already exists" {
		logrus.Error(err)
	}
}
