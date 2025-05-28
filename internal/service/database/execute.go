package database

import (
	"context"
	"fmt"

	"ed_in_memory_database/internal/entity"
)

// Execute -
func (d Database) Execute(ctx context.Context, rawQuery string) (ExecuteResponse, error) {
	command, err := d.compute.ParseQuery(rawQuery)
	if err != nil {
		return ExecuteResponse{
			message: fmt.Sprintf("invalid query: %s", err.Error()),
		}, nil
	}

	switch command.Type() {
	case entity.CommandTypeGet:
		return d.executeGet(ctx, command)
	case entity.CommandTypeSet:
		return d.executeSet(ctx, command)
	case entity.CommandTypeDelete:
		return d.executeDelete(ctx, command)
	default:
	}

	d.logger.Errorf("failed compute query: unknown command type")

	return ExecuteResponse{
		message: internalExecutionFailedMessage,
	}, nil
}
