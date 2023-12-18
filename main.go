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

func main() {
	app := fiber.New()
	app.Get("/api/hello", handleHello)
	app.Get("/api/weather", handleWeather)
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         frontend.BuildHTTPFS(),
		NotFoundFile: "index.html",
	}))
	log.Fatal(app.Listen(fmt.Sprintf(":%s", getEnv("APP_PORT", "4000"))))
}

// Endpoint handlers
func handleHello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello from the server"})
}

func handleWeather(c *fiber.Ctx) error {
	agent := fiber.Get("https://api.weatherapi.com/v1/forecast.json?key=1b71345d3d30453e979123332231412&q=Basing&days=1&aqi=no&alerts=no")

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": errs,
		})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(statusCode).Send(body)
}
