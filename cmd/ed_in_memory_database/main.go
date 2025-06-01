package main

import (
	"context"
	"os"

	"ed_in_memory_database/internal/app"
	"ed_in_memory_database/internal/config"
	core_logger "ed_in_memory_database/pkg/logger"
)

func main() {
	logger, err := core_logger.New(core_logger.Params{
		ServiceName: "ed_in_memory_database",
		LogDir:      "./logs",
		LogLevel:    core_logger.DebugLevel,
	}, os.Stderr)
	if err != nil {
		panic(err)
	}

	cfg := config.Config{}

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatalf("failed to initialize app: %v", err)
	}

	ctx := context.Background()

	err = app.Run(ctx)
	if err != nil {
		logger.Errorf("app failed on run: %v", err)
	} else {
		logger.Errorf("app stopped")
	}
}
