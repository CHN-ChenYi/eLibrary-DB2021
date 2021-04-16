package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	app := fiber.New()

	addRoutes(app)

	port := viper.GetString("app.port")
	app.Listen(":" + port)

	logrus.Info("Echo framework initialized")
}

func addRoutes(app *fiber.App) {
	
}
