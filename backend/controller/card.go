package controller

import (
	"github.com/CHN-ChenYi/eLibrary-DB2021/model"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func validateCard(card *model.Card) []*errorResponse {
	var errors []*errorResponse
	validate := validator.New()
	err := validate.Struct(card)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func modifyCardTemplate(c *fiber.Ctx, modifyCard func(*model.Card) error) error {
	card := new(model.Card)

	if err := c.BodyParser(card); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	if err := validateCard(card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": err,
		})
	}

	if err := modifyCard(card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "SUCCESS"})
}

func getCard(c *fiber.Ctx) error {
	cardID := c.Query("card_id")
	if cardID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "card_id can't be empty",
		})
	}
	cards, err := model.QueryCardByCardID(cardID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	if len(cards) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "can't find such card",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": cards[0]})
}

func addCard(c *fiber.Ctx) error {
	return modifyCardTemplate(c, model.CreateCard)
}

func modifyCard(c *fiber.Ctx) error {
	return modifyCardTemplate(c, model.ModifyCard)
}

func deleteCard(c *fiber.Ctx) error {
	cardID := c.Query("card_id")
	if cardID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "card_id can't be empty",
		})
	}

	books, err := model.QueryBorrowWithoutReturnDateAll(cardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}
	if len(books) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "hasn't returned all the books",
		})
	}

	err = model.DeleteCard(cardID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "SUCCESS"})
}
