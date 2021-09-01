package controllers

import (
	"fiber-go/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func FetchAllPegawai(c *fiber.Ctx) error {
	result, err := models.FethAllPegawai()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(result)
}
