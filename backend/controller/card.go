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

func addCard(c *fiber.Ctx) error {
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

	if err := model.CreateCard(card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "SUCCESS"})
}

func deleteCard(c *fiber.Ctx) error {
	cardID := c.Query("card_id")
	if cardID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "card_id can't be empty",
		})
	}
	err := model.DeleteCard(cardID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "SUCCESS"})
}
