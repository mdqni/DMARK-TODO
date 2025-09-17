package main

import (
	"context"
	"log"

	"github.com/mdqni/dmark-todo/internal/app"
	"github.com/mdqni/dmark-todo/internal/config"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	newApp, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	if err = newApp.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
