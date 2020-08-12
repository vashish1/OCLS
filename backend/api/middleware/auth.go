package middleware

import (
	"strings"

	"github.com/gofiber/fiber"
	"github.com/vashish1/OnlineClassPortal/backend/api/utility"
	"github.com/vashish1/OnlineClassPortal/backend/pkg/database/student"
	"github.com/vashish1/OnlineClassPortal/backend/pkg/database/teacher"
	"github.com/vashish1/OnlineClassPortal/backend/pkg/models"
)

func Auth() func(c *fiber.Ctx) {
	return (func(c *fiber.Ctx) {
		str := c.Get("Authorization")
		c.Set("Content-Type", "application/json")

		tokenString := strings.TrimPrefix(str, "Bearer ")

		if tokenString != "" {
			m, ok := utility.VerifyJwt(str)
			if !ok {
				resp := models.LoginResponse{
					Success: false,
					Token:   "",
					Error:   "Invalid Token",
				}
				c.Status(401).JSON(resp)
				return
			}
			userType := m["type"].(float64)
			if userType == 0 {
				ok = student.IsAvailable(m["uid"].(string))
			} else {
				ok = teacher.IsAvailable(m["uid"].(string))
			}
			if !ok {
				resp := models.LoginResponse{
					Success: false,
					Token:   "",
					Error:   "Authentication failed",
				}
				c.Status(401).JSON(resp)
				return
			}
			c.Next()
		} else {
			resp := models.LoginResponse{
				Success: false,
				Token:   "",
				Error:   "Authorization token required",
			}
			c.Status(401).JSON(resp)
			return
		}
	})
}
