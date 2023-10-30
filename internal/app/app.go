package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"sample/internal/config"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	ctx      context.Context
	cancel   context.CancelFunc
	cfg      *config.Config
	stopChan chan struct{}
	server   *fiber.App
}

func New(cfg *config.Config, f *fiber.App) *App {
	return &App{
		server: f,
		cfg:    cfg,
	}
}

func (a *App) Run() {
	a.ctx, a.cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer a.cancel()
	a.stopChan = make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		select {
		case <-sigint:
			log.Printf("shutdown sequence initiated")

		case <-a.ctx.Done():
			log.Printf("context sequence initiated")

		}
		a.cancel() // руками и ДО шатдауна
		if err := a.server.Shutdown(); err != nil {
			log.Printf("server is not shutting: %v", err)
		}

		close(a.stopChan)
	}()

	if err := a.server.Listen(fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port)); err != nil {
		log.Printf("server is not running: %v", err)
	}

	<-a.stopChan
	log.Printf("shutdown complete")
}

func (a *App) Stop() {
	a.cancel()
}
