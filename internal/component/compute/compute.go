package compute

import (
	"ed_in_memory_database/internal/entity"
)

var (
	queryToCommand = map[string]entity.CommandType{
		"GET": entity.CommandTypeGet,
		"SET": entity.CommandTypeSet,
		"DEL": entity.CommandTypeDelete,
	}
)

// Compute -
type Compute struct {
}

// Make -
func Make() Compute {
	return Compute{}
}
