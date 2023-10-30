package app_test

import (
	"testing"
	"time"

	"sample/internal/app"
	"sample/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestApp(t *testing.T) {
	defer goleak.VerifyNone(t)
	cfg := fiber.Config{
		//	ReadTimeout: 1 * time.Second,
	}
	appCfg := config.New("127.0.0.1", 8080)

	app := app.New(appCfg, fiber.New(cfg))

	go app.Run()
	assert.Equal(t, 1, 1)
	time.Sleep(time.Second * 3)
	app.Stop()
	time.Sleep(time.Second * 3)
}

func TestApp2(t *testing.T) {
	defer goleak.VerifyNone(t)
	app := fiber.New() // create a new Fiber instance

	// Create a new endpoint
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World!") // send text
	//})

	// Start server on port 3000
	go app.Listen(":3000")
	time.Sleep(time.Second * 3)
	app.Shutdown()
	time.Sleep(time.Second * 3)
}
