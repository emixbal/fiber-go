package controllers

import (
	"fiber-go/helpers"
	"fiber-go/models"
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
	username := c.FormValue("username")
	txtUnHashPassword := c.FormValue("password")

	isMatch, _ := models.CheckLogin(username, txtUnHashPassword)
	if !isMatch {
		return c.SendString("password salah")
	}
	return c.SendString("ok")
}
