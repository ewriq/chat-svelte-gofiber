package Handlers

import (
	"fmt"
    "strings"
	"chat-app/Database"
	"chat-app/Utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var reqbody UserBody

	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	email := reqbody.Email
	password := reqbody.Password
	fmt.Print(email)
	fmt.Print(password)

	if isValidEmail(email) {
		if isPasswordValid(password) {
			err, token  := Database.Login(email, password)
			fmt.Println(token,err,"----------------------------------------------------")
			if err == "" {
				c.JSON(fiber.Map{
					"status": "ERROR",
					"msg":    "Kullanıcı bulunamadı veya e-posta, şifreniz hatalı ",
				})
			} else if err == "." {
				c.JSON(fiber.Map{
					"status": "ERROR",
					"msg":    "Kullanıcı bulunamadı lütfen kayıt olunuz",
				})
			} else {
				c.JSON(fiber.Map{
					"status": "OK",
					"msg":    "Kullanıcı başarıyla giriş yaptı",
					"token":  err,
				})
			}
		} else {
			c.JSON(fiber.Map{
				"status": "ERROR",
				"msg":    "Geçersiz parola.",
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": "ERROR",
			"msg":    "Geçersiz email.",
		})
	}

	return nil
}
func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func isPasswordValid(password string) bool {
	return len(password) >= 8
}



func Register(c *fiber.Ctx) error {
	var reqbody UserBody

	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	email := reqbody.Email
	password := reqbody.Password
	username := reqbody.Username

	fmt.Print(email)
	fmt.Print(password)
	fmt.Print(username)
	
	if isValidEmail(email) {
		if isPasswordValid(password) {
            token := Utils.Token(10)
			fmt.Print(token)
			err := Database.Register(username, password, email, token)
			if err != true {
				c.JSON(fiber.Map{
					"status": 502,
					"error":  err,
				})
			} else {
				c.JSON(fiber.Map{
					"status": "OK",
					"message": "User registered successfully",
					"token": token,
				})
			}
		} else {
			c.JSON(fiber.Map{
				"status": "ERROR",
				"error":  "Geçersiz parola.",
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": "ERROR",
			"error":  "Geçersiz email.",
		})
	}

	return nil
}

func User(c *fiber.Ctx) error {
	var reqbody UserInfo
	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	token := reqbody.Token
	user, err := Database.Users(token)
	fmt.Println(token)
    fmt.Println("Handler",user)
	if err != nil {
		c.JSON(fiber.Map{
			"status": "error",
		})
		return err
	}

	c.JSON(fiber.Map{
		"status": "OK",
		"data":   user,
	})
	

	return nil
}