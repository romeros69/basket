package app

import (
	"fmt"
	"github.com/romeros69/basket/internal/usecase"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/config"
	v1 "github.com/romeros69/basket/internal/controller/http/v1"
	"github.com/romeros69/basket/pkg/httpserver"
	"github.com/romeros69/basket/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Use case
	translationUseCase := usecase.NewHelloWorld()

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, translationUseCase, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	l.Info("server is start")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		}
	}
}
