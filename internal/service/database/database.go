package database

import (
	"context"
	"errors"

	"ed_in_memory_database/internal/entity"
	"ed_in_memory_database/internal/entity/storing_types"
	"ed_in_memory_database/pkg/logger"
)

type (
	computeLayer interface {
		ParseQuery(rawQuery string) (entity.Command, error)
	}

	storageLayer interface {
		Get(ctx context.Context, key storing_types.KeyStr) (storing_types.ValueStr, error)
		Set(ctx context.Context, key storing_types.KeyStr, value storing_types.ValueStr) error
		Delete(ctx context.Context, key storing_types.KeyStr) error
	}
)

// Database -
type Database struct {
	storage storageLayer
	compute computeLayer
	logger  logger.Logger
}

// MakeDatabase -
func MakeDatabase(storage storageLayer, compute computeLayer, logger logger.Logger) (Database, error) {
	if storage == nil {
		return Database{}, errors.New("database storage is nil")
	}
	if compute == nil {
		return Database{}, errors.New("database compute is nil")
	}
	if logger == nil {
		return Database{}, errors.New("database logger is nil")
	}
	return Database{
		storage: storage,
		compute: compute,
		logger:  logger,
	}, nil
}
