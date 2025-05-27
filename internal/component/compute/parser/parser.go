package parser

import (
	"errors"
	"fmt"
	"strings"

	"ed_in_memory_database/internal/entity"
)

var (
	queryToCommand = map[string]entity.CommandType{
		"GET": entity.CommandTypeGet,
		"SET": entity.CommandTypeSet,
		"DEL": entity.CommandTypeDelete,
	}
)

// ParseQuery -
func ParseQuery(rawQuery string) (entity.Command, error) {
	tokens := strings.Fields(rawQuery)

	if len(tokens) == 0 {
		return entity.Command{}, errors.New("empty query")
	}

	commandStr := tokens[0]

	commandType, ok := queryToCommand[commandStr]
	if !ok {
		return entity.Command{}, fmt.Errorf("\"%s\" unknown command", commandStr)
	}

	return entity.Make(commandType, tokens[1:]...)
}
