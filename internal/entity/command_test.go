package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMake(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		commandType CommandType
		args        []string

		expected    Command
		expectedErr string
	}{
		{
			name:        "get",
			commandType: CommandTypeGet,
			args:        []string{"key"},

			expected: Command{
				commandType: CommandTypeGet,
				args:        []string{"key"},
			},
		},
		{
			name:        "set",
			commandType: CommandTypeSet,
			args:        []string{"key", "val"},

			expected: Command{
				commandType: CommandTypeSet,
				args:        []string{"key", "val"},
			},
		},
		{
			name:        "delete",
			commandType: CommandTypeDelete,
			args:        []string{"key"},
			expected: Command{
				commandType: CommandTypeDelete,
				args:        []string{"key"},
			},
		},
		{
			name:        "invalid_type_command",
			commandType: CommandTypeUnknown,

			expectedErr: "invalid command type",
		},
		{
			name:        "incorrect_args_count",
			commandType: CommandTypeGet,
			args:        []string{"key", "val"},

			expectedErr: "incorrect number of arguments: awaiting 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			command, err := Make(tt.commandType, tt.args...)
			if len(tt.expectedErr) > 0 {
				require.Error(t, err)
				require.Equal(t, tt.expectedErr, err.Error())
			}
			require.Equal(t, tt.expected, command)
		})
	}
}
