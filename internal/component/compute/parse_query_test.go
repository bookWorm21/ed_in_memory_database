package compute

import (
	"testing"

	"ed_in_memory_database/internal/entity"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		query string

		expectedCommand entity.Command
		expectedErr     string
	}{
		{
			name:  "get_command",
			query: "GET val",

			expectedCommand: entity.MakerCommand{
				CommandType: entity.CommandTypeGet,
				Args:        []string{"val"},
			}.Make(t),
		},
		{
			name:  "set_command",
			query: "SET val 5",

			expectedCommand: entity.MakerCommand{
				CommandType: entity.CommandTypeSet,
				Args:        []string{"val", "5"},
			}.Make(t),
		},
		{
			name:  "delete_command",
			query: "DEL val",

			expectedCommand: entity.MakerCommand{
				CommandType: entity.CommandTypeDelete,
				Args:        []string{"val"},
			}.Make(t),
		},
		{
			name:  "invalid_args_count",
			query: "GET val 1",

			expectedErr: "incorrect number of arguments: awaiting 1",
		},
		{
			name:  "invalid_command",
			query: "UPDATE val 15",

			expectedErr: "\"UPDATE\" unknown command",
		},
		{
			name:  "low_case_command",
			query: "Get val",

			expectedErr: "\"Get\" unknown command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			compute := Make()

			command, err := compute.ParseQuery(tt.query)
			if len(tt.expectedErr) > 0 {
				require.Error(t, err)
				require.Equal(t, tt.expectedErr, err.Error())
			}

			require.Equal(t, tt.expectedCommand, command)
		})
	}
}
