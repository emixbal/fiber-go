package controllers

import "github.com/gofiber/fiber/v2"

// func FetchAllPegawai(c echo.Context) error {
// 	result, err := models.FethAllPegawai()

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}
// 	return c.JSON(http.StatusOK, result)
// }

func FetchAllPegawai(c *fiber.Ctx) error {
	return c.SendString("fetch all pegawai")
}
