package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	ErrorDetail string `json:"error"`
}

func HandleError(c *fiber.Ctx, message string, err string, statusCode int) error {
	response := ErrorResponse{
		Success:     false,
		Message:     message,
		ErrorDetail: err,
	}
	logrus.Error(err)
	c.Status(statusCode)
	return c.JSON(response)
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(c *fiber.Ctx, message string, data interface{}) error {
	response := SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
