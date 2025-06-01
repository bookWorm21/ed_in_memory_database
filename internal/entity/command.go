package entity

import (
	"fmt"
	"strconv"
)

// CommandType -
type CommandType int

const (
	// CommandTypeUnknown -
	CommandTypeUnknown CommandType = iota
	// CommandTypeGet -
	CommandTypeGet
	// CommandTypeSet -
	CommandTypeSet
	// CommandTypeDelete -
	CommandTypeDelete
)

// commandTypeToArgsCount - количество аргументов по типу команды необходимых для выполнения
var commandTypeToArgsCount = map[CommandType]int{
	CommandTypeGet:    1,
	CommandTypeSet:    2,
	CommandTypeDelete: 1,
}

var (
	// ErrIncorrectNumberOfArguments -
	ErrIncorrectNumberOfArguments = fmt.Errorf("incorrect number of arguments")
	// ErrInvalidCommandType -
	ErrInvalidCommandType = fmt.Errorf("invalid command type")
)

// Command -
type Command struct {
	commandType CommandType
	args        []string
}

// Make - конструирует Command
func Make(commandType CommandType, args ...string) (Command, error) {
	if commandType == CommandTypeUnknown {
		return Command{}, ErrInvalidCommandType
	}

	needArgsCount := commandTypeToArgsCount[commandType]
	if needArgsCount != len(args) {
		return Command{}, fmt.Errorf("%w: awaiting %s", ErrIncorrectNumberOfArguments, strconv.Itoa(needArgsCount))
	}

	return Command{
		commandType: commandType,
		args:        args,
	}, nil
}

// Type -
func (c Command) Type() CommandType {
	return c.commandType
}

// Arg -
func (c Command) Arg(number int) (string, error) {
	if number < 0 || number >= len(c.args) {
		return "", ErrIncorrectNumberOfArguments
	}
	return c.args[number], nil
}
