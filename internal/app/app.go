package app

import (
	"context"
	"errors"

	"ed_in_memory_database/internal/cli"
	compute_component "ed_in_memory_database/internal/component/compute"
	storage_component "ed_in_memory_database/internal/component/storage"
	"ed_in_memory_database/internal/component/storage/engine/in_memory"
	"ed_in_memory_database/internal/config"
	database_service "ed_in_memory_database/internal/service/database"
	"ed_in_memory_database/pkg/logger"
)

// App -
type App struct {
	logger logger.Logger
	cli    cli.Cli
}

// NewApp -
func NewApp(_ config.Config, logger logger.Logger) (*App, error) {
	engine := in_memory.MakeEngine()
	compute := compute_component.Make()
	storage, err := storage_component.Make(engine)
	if err != nil {
		return nil, errors.New("failed create storage")
	}
	database, err := database_service.MakeDatabase(storage, compute, logger)
	if err != nil {
		return nil, errors.New("failed create database")
	}
	cli, err := cli.Make(database)
	if err != nil {
		return nil, errors.New("failed create cli")
	}
	return &App{
		logger: logger,
		cli:    cli,
	}, nil
}

// Run -
func (a *App) Run(ctx context.Context) error {
	a.cli.Run(ctx)
	return nil
}
