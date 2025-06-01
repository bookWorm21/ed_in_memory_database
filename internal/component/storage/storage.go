package storage

import (
	"context"
	"errors"

	"ed_in_memory_database/internal/entity/storing_types"
)

type (
	engine interface {
		Get(ctx context.Context, key storing_types.KeyStr) (storing_types.ValueStr, error)
		Set(ctx context.Context, key storing_types.KeyStr, value storing_types.ValueStr) error
		Delete(ctx context.Context, key storing_types.KeyStr) error
	}
)

// Storage -
type Storage struct {
	engine engine
}

// Make -
func Make(engine engine) (Storage, error) {
	if engine == nil {
		return Storage{}, errors.New("engine is nil")
	}
	return Storage{
		engine: engine,
	}, nil
}

// Get -
func (s Storage) Get(ctx context.Context, key storing_types.KeyStr) (storing_types.ValueStr, error) {
	return s.engine.Get(ctx, key)
}

// Set -
func (s Storage) Set(ctx context.Context, key storing_types.KeyStr, value storing_types.ValueStr) error {
	return s.engine.Set(ctx, key, value)
}

// Delete -
func (s Storage) Delete(ctx context.Context, key storing_types.KeyStr) error {
	return s.engine.Delete(ctx, key)
}
