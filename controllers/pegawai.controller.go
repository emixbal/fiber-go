package controllers

import "github.com/gofiber/fiber/v2"

func FetchAllPegawai(c *fiber.Ctx) error {
	return c.SendString("fetch all pegawai")
}
