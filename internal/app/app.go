package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/romeros69/basket/config"
	v1 "github.com/romeros69/basket/internal/controller/http/v1"
	"github.com/romeros69/basket/internal/usecase"
	"github.com/romeros69/basket/internal/usecase/repo/mongo_rp"
	"github.com/romeros69/basket/pkg/httpserver"
	"github.com/romeros69/basket/pkg/logger"
	"github.com/romeros69/basket/pkg/mongo"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	//// MongoDB
	mongoDB, err := mongo.New(cfg)
	if err != nil {
		panic(err)
	}

	// Repository
	playerRepo := mongo_rp.NewPlayerRepo(mongoDB, "players")

	// Use case
	playerUseCase := usecase.NewPlayerUC(playerRepo)

	// HTTP Server
	handler := gin.New()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Content-Type", "Access-Control-Allow-Credentials", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v1.NewRouter(handler, playerUseCase, l)
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
