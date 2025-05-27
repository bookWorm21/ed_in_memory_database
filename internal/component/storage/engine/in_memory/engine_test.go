package in_memory

import (
	"context"
	"testing"

	storage_errors "ed_in_memory_database/internal/component/storage/errors"
	"ed_in_memory_database/internal/entity/storing_types"

	"github.com/stretchr/testify/require"
)

func TestEngine_KeyCoreProperties(t *testing.T) {
	t.Parallel()

	key := storing_types.KeyStr("key")
	value := storing_types.ValueStr("value")
	ctx := context.Background()

	engine := MakeEngine()
	_, err := engine.Get(ctx, key)
	require.ErrorIs(t, err, storage_errors.ErrNotFound)

	err = engine.Set(ctx, key, value)
	require.NoError(t, err)

	val, err := engine.Get(ctx, key)
	require.NoError(t, err)
	require.Equal(t, value, val)

	err = engine.Delete(ctx, key)
	require.NoError(t, err)

	_, err = engine.Get(ctx, key)
	require.ErrorIs(t, err, storage_errors.ErrNotFound)
}

func TestEngine_Get_Idempotence(t *testing.T) {
	t.Parallel()

	key := storing_types.KeyStr("key")
	value := storing_types.ValueStr("value")
	ctx := context.Background()

	engine := MakeEngine()
	err := engine.Set(ctx, key, value)
	require.NoError(t, err)

	val, err := engine.Get(ctx, key)
	require.NoError(t, err)
	require.Equal(t, value, val)

	val, err = engine.Get(ctx, key)
	require.NoError(t, err)
	require.Equal(t, value, val)
}

func TestEngine_Delete_Idempotence(t *testing.T) {
	t.Parallel()

	key := storing_types.KeyStr("key")
	value := storing_types.ValueStr("value")
	ctx := context.Background()

	engine := MakeEngine()
	err := engine.Set(ctx, key, value)
	require.NoError(t, err)

	err = engine.Delete(ctx, key)
	require.NoError(t, err)

	err = engine.Delete(ctx, key)
	require.NoError(t, err)
}

func TestEngine_Set_ChangedValue(t *testing.T) {
	t.Parallel()

	key := storing_types.KeyStr("key")
	value := storing_types.ValueStr("value")
	value2 := storing_types.ValueStr("value2")
	ctx := context.Background()

	engine := MakeEngine()
	err := engine.Set(ctx, key, value)
	require.NoError(t, err)

	val, err := engine.Get(ctx, key)
	require.NoError(t, err)
	require.Equal(t, value, val)

	err = engine.Set(ctx, key, value2)
	require.NoError(t, err)

	val, err = engine.Get(ctx, key)
	require.NoError(t, err)
	require.Equal(t, value2, val)
}
