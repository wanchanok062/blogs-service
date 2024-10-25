package middleware

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/wanchanok6698/web-blogs/util"
)

var validation *validator.Validate

func init() {
	validation = validator.New()
}

func validateStruct(data interface{}) error {
	return validation.Struct(data)
}

func ValidateData(v interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(v); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}
		if err := validateStruct(v); err != nil {
			log.Println(err.Error())
			return util.HandleError(c, "Validation error", err.Error(), fiber.StatusBadRequest)
		}
		return c.Next()
	}
}
