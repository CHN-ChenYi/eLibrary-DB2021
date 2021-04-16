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
		logrus.Fatal("DB is nil")
	}
	err := gormDb.AutoMigrate(
		&Book{},
		&Card{},
		&Borrow{},
	)
	if err != nil {
		logrus.Fatal(err)
	}
}
