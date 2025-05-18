package database

import (
	"context"
	"errors"
	"fmt"

	storage_errors "ed_in_memory_database/internal/component/storage/errors"
	"ed_in_memory_database/internal/entity"
	"ed_in_memory_database/internal/entity/storing_types"
)

func (d Database) executeGet(ctx context.Context, command entity.Command) (ExecuteResponse, error) {
	key, err := command.Arg(0)
	if err != nil {
		d.logger.Errorf("failed get arg: %s", err.Error())
		return ExecuteResponse{
			message: internalExecutionFailedMessage,
		}, nil
	}

	val, err := d.storage.Get(ctx, storing_types.KeyStr(key))

	if errors.Is(err, storage_errors.ErrNotFound) {
		return ExecuteResponse{
			message: fmt.Sprintf("key %s does not exist", key),
		}, nil
	}

	if err != nil {
		d.logger.Errorf("executing get command: %s", err.Error())
		return ExecuteResponse{
			message: internalExecutionFailedMessage,
		}, nil
	}

	return ExecuteResponse{
		message: fmt.Sprintf("value: %s", val),
	}, nil
}

func (d Database) executeSet(ctx context.Context, command entity.Command) (ExecuteResponse, error) {
	key, err := command.Arg(0)
	if err != nil {
		d.logger.Errorf("failed get arg: %s", err.Error())
		return ExecuteResponse{
			message: internalExecutionFailedMessage,
		}, nil
	}
	value, err := command.Arg(1)
	if err != nil {
		d.logger.Errorf("failed get arg: %s", err.Error())
		return ExecuteResponse{
			message: internalExecutionFailedMessage,
		}, nil
	}

	err = d.storage.Set(ctx, storing_types.KeyStr(key), storing_types.ValueStr(value))

	if err != nil {
		d.logger.Errorf("executing set command: %s", err.Error())
		return ExecuteResponse{
			message: internalExecutionFailedMessage,
		}, nil
	}

	return ExecuteResponse{
		message: "success set",
	}, nil
}

func (d Database) executeDelete(ctx context.Context, command entity.Command) (ExecuteResponse, error) {
	key, err := command.Arg(0)
	if err != nil {
		d.logger.Errorf("failed get arg: %s", err.Error())
		return ExecuteResponse{
			message: internalExecutionFailedMessage,
		}, nil
	}

	err = d.storage.Delete(ctx, storing_types.KeyStr(key))
	if err != nil {
		d.logger.Errorf("executing delete command: %s", err.Error())
		return ExecuteResponse{
			message: internalExecutionFailedMessage,
		}, nil
	}

	return ExecuteResponse{
		message: "success delete",
	}, nil
}
