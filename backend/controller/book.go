package controller

import (
	"strconv"

	"github.com/CHN-ChenYi/eLibrary-DB2021/model"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func validateBook(book *model.Book) []*errorResponse {
	var errors []*errorResponse
	validate := validator.New()
	err := validate.Struct(book)
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

func addBook(c *fiber.Ctx) error {
	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	if err := validateBook(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": err,
		})
	}

	if err := model.CreateBook(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "SUCCESS"})
}

func modifyBook(c *fiber.Ctx) error {
	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	if err := validateBook(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": err,
		})
	}

	if err := model.ModifyBook(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "SUCCESS"})
}

func getBookAll(c *fiber.Ctx) error {
	books, err := model.QueryBookAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": books})
}

func getBookBySomething(field string, QueryBookBySomething func(string) ([]model.Book, error)) (func(c *fiber.Ctx) error) {
	return func(c *fiber.Ctx) error {
		something := c.Query(field)
		if something == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": field + " can't be empty",
			})
		}
		books, err := QueryBookBySomething(something)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"data": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": books})
	}
}

func getBookByYear(c *fiber.Ctx) error {
	l_str := c.Query("l")
	r_str := c.Query("r")
	l, l_err := strconv.Atoi(l_str)
	r, r_err := strconv.Atoi(r_str)
	if l_err != nil || r_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "couldn't parse range",
		})
	}
	books, err := model.QueryBookByYear(l, r)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": books})
}

func getBookByPrice(c *fiber.Ctx) error {
	l_str := c.Query("l")
	r_str := c.Query("r")
	l_double, l_err := strconv.ParseFloat(l_str, 32)
	r_double, r_err := strconv.ParseFloat(r_str, 32)
	if l_err != nil || r_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "couldn't parse range",
		})
	}
	l := float32(l_double)
	r := float32(r_double)
	books, err := model.QueryBookByPrice(l, r)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": books})
}
