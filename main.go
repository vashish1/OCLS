package main

import (
	"os"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	mw "github.com/vashish1/OnlineClassPortal/api/middleware"
	v1 "github.com/vashish1/OnlineClassPortal/api/v1"
)

var router *fiber.App
var login, dash fiber.Router
var port = os.Getenv("PORT")

func setupRoutes() {
	login.Post("/student", v1.StudentsLogin)
	login.Post("/teacher", v1.TeachersLogin)
	dash.Get("/", v1.Dashboard)
	router.Get("/ws",v1.ServeWs)
	// router.Post("/signup",v1.Sign)

}

func main() {
	// if port == "" {
	// 	port = "8000"
	// }
	router = fiber.New()
	router.Use(cors.New())
	login = router.Group("/api/login")
	dash = router.Group("/api/dashboard", mw.Auth())

	setupRoutes()
	router.Listen(port)
}
