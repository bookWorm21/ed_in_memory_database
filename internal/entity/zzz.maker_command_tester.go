package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type MakerCommand struct {
	CommandType CommandType
	Args        []string
}

func (m MakerCommand) Make(t *testing.T) Command {
	t.Helper()

	command, err := Make(m.CommandType, m.Args...)
	require.NoError(t, err)

	return command
}
