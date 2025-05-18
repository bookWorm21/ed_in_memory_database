package app

import (
	"context"

	"ed_in_memory_database/internal/config"
	"ed_in_memory_database/pkg/logger"
)

type App struct {
	logger logger.Logger
}

func NewApp(_ config.Config, logger logger.Logger) (*App, error) {
	return &App{
		logger: logger,
	}, nil
}

func (a *App) Run(_ context.Context) error {
	return nil
}
