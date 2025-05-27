package entity

import (
	"fmt"
	"strconv"
)

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

// CommandTypeToArgsCount - количество аргументов по типу команды необходимых для выполнения
var CommandTypeToArgsCount = map[CommandType]int{
	CommandTypeGet:    1,
	CommandTypeSet:    2,
	CommandTypeDelete: 1,
}

var (
	ErrIncorrectNumberOfArguments = fmt.Errorf("incorrect number of arguments")
	ErrInvalidCommandType         = fmt.Errorf("invalid command type")
)

type Command struct {
	commandType CommandType
	args        []string
}

// Make - конструирует Command
func Make(commandType CommandType, args ...string) (Command, error) {
	if commandType == CommandTypeUnknown {
		return Command{}, ErrInvalidCommandType
	}

	needArgsCount := CommandTypeToArgsCount[commandType]
	if needArgsCount != len(args) {
		return Command{}, fmt.Errorf("%w: awaiting %s", ErrIncorrectNumberOfArguments, strconv.Itoa(needArgsCount))
	}

	return Command{
		commandType: commandType,
		args:        args,
	}, nil
}

func (c Command) Type() CommandType {
	return c.commandType
}

func (c Command) Arg(number int) (string, error) {
	if number < 0 || number >= len(c.args) {
		return "", ErrIncorrectNumberOfArguments
	}
	return c.args[number], nil
}
