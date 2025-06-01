package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// MakerCommand - creating Command on tests
type MakerCommand struct {
	CommandType CommandType
	Args        []string
}

// Make -
func (m MakerCommand) Make(t *testing.T) Command {
	t.Helper()

	command, err := Make(m.CommandType, m.Args...)
	require.NoError(t, err)

	return command
}
