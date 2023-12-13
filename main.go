package main

import (
	"fmt"
	"log"
	"os"
	"svelte-go-starter/frontend"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func setupRoutes(app *fiber.App) {

	// Public routes
	app.Get("/api/hello", handleHello)
}

func main() {
	app := fiber.New()
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         frontend.BuildHTTPFS(),
		NotFoundFile: "index.html",
	}))
	setupRoutes(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", getEnv("APP_PORT", "4000"))))
}

// Endpoint handlers
func handleHello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello from the server"})
}
