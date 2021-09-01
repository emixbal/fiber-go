package controllers

import (
	"fiber-go/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FetchAllPegawais(c *fiber.Ctx) error {
	result, err := models.FethAllPegawai()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(result)
}

func UpdatePegawai(c *fiber.Ctx) error {
	id := c.Params("id")
	name := c.FormValue("name")
	alamat := c.FormValue("alamat")
	telephone := c.FormValue("telephone")
	converted_id, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	result, err := models.UpdatePegawai(converted_id, name, alamat, telephone)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(result)
}
