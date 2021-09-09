package controllers

import (
	"fiber-go/config"
	"fiber-go/helpers"
	"fiber-go/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type jwtClaims struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

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

	conf := config.GetConfig()

	isMatch, _ := models.CheckLogin(username, txtUnHashPassword)
	if !isMatch {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "Password salah"})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(conf.JWT_SECRET))
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "something went wrong!"})
	}
	return c.Status(http.StatusOK).JSON(map[string]string{
		"token": tokenString,
	})
}
