package controller

import (
	"github.com/CHN-ChenYi/eLibrary-DB2021/model"
	"github.com/gofiber/fiber/v2"
)

func validateCardID(c *fiber.Ctx) error {
	cardID := c.Query("card_id")
	if cardID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "card_id can't be empty",
		})
	}
	exists, err := model.ValidateCard(cardID)
	if !exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "card_id invalid",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Next()
}

func getBorrowWithoutReturnDateAll(c *fiber.Ctx) error {
	cardID := c.Query("card_id")

	books, err := model.QueryBorrowWithoutReturnDateAll(cardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": books})
}

func modifyBorrowTemplate(c *fiber.Ctx, modifyBorrow func(string, string) error) error {
	cardID := c.Query("card_id")
	bookID := c.Query("book_id")
	if bookID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "book_id can't be empty",
		})
	}

	err := modifyBorrow(cardID, bookID)
	if err != nil {
		if err == model.ErrNoRowsAffected {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": "no such borrow record",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "SUCCESS"})
}

func borrowBook(c *fiber.Ctx) error {
	return modifyBorrowTemplate(c, model.CreateBorrow)
}

func returnBook(c *fiber.Ctx) error {
	return modifyBorrowTemplate(c, model.DeleteBorrow)
}
