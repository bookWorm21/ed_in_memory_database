package database

import (
	"context"
	"errors"
	"fmt"

	storage_errors "ed_in_memory_database/internal/component/storage/errors"
	"ed_in_memory_database/internal/entity"
	"ed_in_memory_database/internal/entity/storing_types"
)

const (
	failedArgErrMessage        = "failed get args: %s"
	executingCommandErrMessage = "%s executing command failed: %s"
)

func (d Database) executeGet(ctx context.Context, command entity.Command) (ExecuteResponse, error) {
	key, err := command.Arg(0)
	if err != nil {
		return d.internalExecutionFailedResponseWithLogError(fmt.Sprintf(failedArgErrMessage, err.Error())), nil
	}

	val, err := d.storage.Get(ctx, storing_types.KeyStr(key))

	if errors.Is(err, storage_errors.ErrNotFound) {
		return ExecuteResponse{
			message: fmt.Sprintf("key \"%s\" does not exist", key),
		}, nil
	}

	if err != nil {
		return d.internalExecutionFailedResponseWithLogError(fmt.Sprintf(executingCommandErrMessage, "get", err.Error())), nil
	}

	return ExecuteResponse{
		message: fmt.Sprintf("value: \"%s\"", val),
	}, nil
}

func (d Database) executeSet(ctx context.Context, command entity.Command) (ExecuteResponse, error) {
	key, err := command.Arg(0)
	if err != nil {
		return d.internalExecutionFailedResponseWithLogError(fmt.Sprintf(failedArgErrMessage, err.Error())), nil
	}
	value, err := command.Arg(1)
	if err != nil {
		return d.internalExecutionFailedResponseWithLogError(fmt.Sprintf(failedArgErrMessage, err.Error())), nil
	}

	err = d.storage.Set(ctx, storing_types.KeyStr(key), storing_types.ValueStr(value))

	if err != nil {
		return d.internalExecutionFailedResponseWithLogError(fmt.Sprintf(executingCommandErrMessage, "set", err.Error())), nil
	}

	return ExecuteResponse{
		message: "success set",
	}, nil
}

func (d Database) executeDelete(ctx context.Context, command entity.Command) (ExecuteResponse, error) {
	key, err := command.Arg(0)
	if err != nil {
		return d.internalExecutionFailedResponseWithLogError(fmt.Sprintf(failedArgErrMessage, err.Error())), nil
	}

	err = d.storage.Delete(ctx, storing_types.KeyStr(key))
	if err != nil {
		return d.internalExecutionFailedResponseWithLogError(fmt.Sprintf(executingCommandErrMessage, "delete", err.Error())), nil
	}

	return ExecuteResponse{
		message: "success delete",
	}, nil
}

func (d Database) internalExecutionFailedResponseWithLogError(errMessage string) ExecuteResponse {
	d.logger.Errorf(errMessage)
	return ExecuteResponse{
		message: internalExecutionFailedMessage,
	}
}
