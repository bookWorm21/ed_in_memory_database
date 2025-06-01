package in_memory

import (
	"context"

	storage_errors "ed_in_memory_database/internal/component/storage/errors"
	"ed_in_memory_database/internal/entity/storing_types"
)

// Engine  -
type Engine struct {
	data ConcurrentHashTable[storing_types.KeyStr, storing_types.ValueStr]
}

// MakeEngine -
func MakeEngine() Engine {
	return Engine{
		data: MakeConcurrentHashTable[storing_types.KeyStr, storing_types.ValueStr](),
	}
}

// Get -
func (s Engine) Get(_ context.Context, key storing_types.KeyStr) (storing_types.ValueStr, error) {
	val, found := s.data.Get(key)
	if !found {
		return "", storage_errors.ErrNotFound
	}
	return val, nil
}

// Set -
func (s Engine) Set(_ context.Context, key storing_types.KeyStr, value storing_types.ValueStr) error {
	s.data.Set(key, value)
	return nil
}

// Delete -
func (s Engine) Delete(_ context.Context, key storing_types.KeyStr) error {
	s.data.Delete(key)
	return nil
}
