package main

import (
	"time"

	"sample/internal/app"
	"sample/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := fiber.Config{
		ReadTimeout: 10 * time.Second,
	}
	appCfg := config.New("127.0.0.1", 8080)
	app := app.New(appCfg, fiber.New(cfg))

	app.Run()
}
