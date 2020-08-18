package v1

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/vashish1/OnlineClassPortal/api/utility"
	"github.com/vashish1/OnlineClassPortal/pkg/models"
)

func StudentsLogin(c *fiber.Ctx) {
	c.Set("Content-Type", "application/json")
	var login models.Login
	var t models.LoginResponse

	err := c.BodyParser(&login)
	fmt.Println("user trying to login \n", login)
	if err != nil {
		t.Success = false
		t.Error = "Body not parsed"
		c.Status(400).JSON(t)
		return
	}

	token, err := utility.GenerateJwtForStudent(login.Email, login.Password)
	if err != nil {
		t.Success = false
		t.Error = err.Error()
		c.Status(400).JSON(t)
		return
	}
	t.Success = true
	t.Token = token
	c.Status(200).JSON(t)
}

func StudentDasboard(c *fiber.Ctx){
	c.Set("Content-Type", "application/json")
	var req map[string]interface{}
	err := c.BodyParser(&req)
	if err != nil {
		 "Body not parsed"
		c.Status(400).JSON(t)
		return
	}
}