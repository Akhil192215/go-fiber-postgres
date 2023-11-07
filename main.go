package main

import (
	"log"
	"os"

	"github.com/Akhil192215/go-fiber/config"
	"github.com/Akhil192215/go-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my awesome api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/user", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/user/:id", routes.GetUser)
	app.Put("/api/user/:id", routes.UpdateUser)
	app.Delete("/api/user/:id", routes.DeleteUser)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	config.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":" + port))
}
