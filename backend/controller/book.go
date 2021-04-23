package controller

import (
	"fmt"
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

func modifyBookTemplate(c *fiber.Ctx, modifyBook func(*model.Book) error) error {
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

	if err := modifyBook(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "SUCCESS"})
}

func addBook(c *fiber.Ctx) error {
	return modifyBookTemplate(c, model.CreateBook)
}

func modifyBook(c *fiber.Ctx) error {
	return modifyBookTemplate(c, model.ModifyBook)
}

func getBook(c *fiber.Ctx) error {
	bookID := c.Query("book_id")
	if bookID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "book_id can't be empty",
		})
	}
	books, err := model.QueryBookByBookID(bookID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	if len(books) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "can't find such book",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": books[0]})
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

func appendQueryString(lhs *string, rhs string) {
	if len(*lhs) == 0 {
		*lhs = rhs
	} else {
		*lhs = *lhs + " AND " + rhs
	}
}

// TODO(TO/GA): prevent XSS
func searchBook(c *fiber.Ctx) error {
	var queryString string

	if category := c.Query("category"); category != "" {
		appendQueryString(&queryString, fmt.Sprintf("category = '%v'", category))
	}

	if title := c.Query("title"); title != "" {
		appendQueryString(&queryString, fmt.Sprintf("title = '%v'", title))
	}

	if press := c.Query("press"); press != "" {
		appendQueryString(&queryString, fmt.Sprintf("press = '%v'", press))
	}

	if author := c.Query("author"); author != "" {
		appendQueryString(&queryString, fmt.Sprintf("author = '%v'", author))
	}

	year_l_str := c.Query("year_lowerbound")
	year_r_str := c.Query("year_upperbound")
	if year_l_str != "" || year_r_str != "" {
		if year_l_str == "" || year_r_str == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": "year_lowerbound and year_upperbound should submit together",
			})
		}
		year_l, l_err := strconv.Atoi(year_l_str)
		year_r, r_err := strconv.Atoi(year_r_str)
		if l_err != nil || r_err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": "couldn't parse range",
			})
		}
		appendQueryString(&queryString, fmt.Sprintf("%v <= year AND year <= %v", year_l, year_r))
	}

	price_l_str := c.Query("price_lowerbound")
	price_r_str := c.Query("price_upperbound")
	if price_l_str != "" || price_r_str != "" {
		if price_l_str == "" || price_r_str == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": "price_lowerbound and price_upperbound should submit together",
			})
		}
		price_l, l_err := strconv.ParseFloat(price_l_str, 32)
		price_r, r_err := strconv.ParseFloat(price_r_str, 32)
		if l_err != nil || r_err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": "couldn't parse range",
			})
		}
		appendQueryString(&queryString, fmt.Sprintf("%v <= price AND price <= %v", price_l, price_r))
	}

	books, err := model.QueryBook(queryString)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": books})
}
