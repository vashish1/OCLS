package v1

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/vashish1/OnlineClassPortal/api/utility"
	"github.com/vashish1/OnlineClassPortal/pkg/database/student"
	"github.com/vashish1/OnlineClassPortal/pkg/database/teacher"
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

// func StudentDasboard(c *fiber.Ctx){
// 	c.Set("Content-Type", "application/json")
// 	var res models.LoginResponse
// 	id:=c.Locals("uid")
// 	data:=student.GetStudent(id.(string))
// 	if data.Email!=""{
// 		c.Status(200).JSON(data)
// 		return
// 	}
// 	res.Success=false
// 	res.Error="Invalid request"
// 	c.Status(400).JSON(res)
// 	return
// }

func Dasboard(c *fiber.Ctx) {
	c.Set("Content-Type", "application/json")
	var res models.LoginResponse
	var data interface{}
	id := c.Locals("uid")
	typ := c.Locals("type")
    
	if typ.(float64)== 0 {
		data = student.Get(id.(string))
	} else {
		data = teacher.Get(id.(string))
	}
	fmt.Println(data)
	if data != nil {
		c.Status(200).JSON(data)
		return
	}
	res.Success = false
	res.Error = "Invalid request"
	c.Status(400).JSON(res)
	return
}
