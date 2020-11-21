package v1

import (
	"github.com/gofiber/fiber"
	"github.com/vashish1/OnlineClassPortal/pkg/database/student"
	"github.com/vashish1/OnlineClassPortal/pkg/database/teacher"
	"github.com/vashish1/OnlineClassPortal/pkg/models"
)

//Dashboard returns the details of user.
func Dashboard(c *fiber.Ctx) {
	c.Set("Content-Type", "application/json")
	var res models.LoginResponse
	var data interface{}
	id := fn0(c)
	typ := fn1(c)

	// fmt.Println("kjhg", typ, "\n")

	if typ.(float64) == 0 {
		data = student.Get(id.(string))
	} else {
		data = teacher.Get(id.(string))
	}
	// fmt.Println(data)
	if data != nil {
		c.Status(200).JSON(data)
		return
	}
	res.Success = false
	res.Error = "Invalid request"
	c.Status(400).JSON(res)
	return
}

func fn1(c *fiber.Ctx) interface{} {
	typ := c.Locals("type")
	return typ
}

func fn0(c *fiber.Ctx) interface{} {
	id := c.Locals("uid")
	return id
}