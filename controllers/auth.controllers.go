package controllers

import (
	"fiber-go/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	rawPassword := c.FormValue("password")
	hashPassword, err := helpers.GeneratePassword(rawPassword)
	if err != nil {
		fmt.Println(err)
	}
	return c.SendString(hashPassword)
}

func Login(c *fiber.Ctx) error {
	unHashPassword := c.FormValue("password")
	hashPasswordDb := "$2a$10$VQA.gpUGleX0vqkSXJcTjuwtZzUCXVUfuJxoHH4T0BUXZ83pFyLou"
	isMatch, _ := helpers.CheckPaaword(hashPasswordDb, unHashPassword)
	if !isMatch {
		return c.SendString("password salah")
	}
	return c.SendString("ok")
}
