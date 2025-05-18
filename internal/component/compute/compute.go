package compute

import (
	"ed_in_memory_database/internal/component/compute/parser"
	"ed_in_memory_database/internal/entity"
)

// Compute -
type Compute struct {
}

// Make -
func Make() Compute {
	return Compute{}
}

// ParseQuery -
func (c Compute) ParseQuery(rawQuery string) (entity.Command, error) {
	return parser.ParseQuery(rawQuery)
}
