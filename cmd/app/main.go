package main

import (
	"log"

	"github.com/romeros69/basket/config"
	"github.com/romeros69/basket/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
