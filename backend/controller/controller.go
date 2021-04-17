package controller

import (
	"github.com/CHN-ChenYi/eLibrary-DB2021/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type errorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func Init() {
	app := fiber.New()

	addRoutes(app)

	port := viper.GetString("app.port")
	app.Listen(":" + port)

	logrus.Info("Echo framework initialized")
}

func addRoutes(app *fiber.App) {
	api := app.Group("/api")

	book := api.Group("/book")
	book.Post("", addBook)
	book.Put("", modifyBook)
	book.Get("/all", getBookAll)
	book.Get("/category", getBookBySomething("category", model.QueryBookByCategory))
	book.Get("/title", getBookBySomething("title", model.QueryBookByTitle))
	book.Get("/press", getBookBySomething("press", model.QueryBookByPress))
	book.Get("/author", getBookBySomething("author", model.QueryBookByAuthor))
	book.Get("/year", getBookByYear)
	book.Get("/price", getBookByPrice)

	card := api.Group("/card")
	card.Post("", addCard)
	card.Put("", modifyCard)
	card.Delete("", deleteCard)

	borrow := api.Group("/borrow")
	borrow.Use(validateCardID)
	borrow.Get("/book/all", getBorrowWithoutReturnDateAll)
	borrow.Post("/book", borrowBook)
	borrow.Delete("/book", returnBook)
}
