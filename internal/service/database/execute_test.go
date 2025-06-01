package database

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"ed_in_memory_database/internal/entity"
	"ed_in_memory_database/internal/pkg/mocks/logger_mock"
	"ed_in_memory_database/internal/service/database/mocks"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDatabase_Execute(t *testing.T) {
	t.Parallel()

	errorMessage := gofakeit.Word()

	tests := []struct {
		name     string
		rawQuery string

		mockCompute func(computeLayer *mocks.ComputeLayerMock)
		mockStorage func(storageLayer *mocks.StorageLayerMock)

		expectedResp ExecuteResponse
		expectedErr  string
	}{
		{
			name:     "get_query",
			rawQuery: "GET val",

			mockCompute: func(computeLayer *mocks.ComputeLayerMock) {
				computeLayer.ParseQueryMock.Times(1).
					Expect("GET val").
					Return(entity.MakerCommand{
						CommandType: entity.CommandTypeGet,
						Args:        []string{"val"},
					}.Make(t), nil)
			},
			mockStorage: func(storageLayer *mocks.StorageLayerMock) {
				storageLayer.GetMock.Times(1).
					Expect(minimock.AnyContext, "val").
					Return("5", nil)
			},

			expectedResp: ExecuteResponse{
				message: "value: \"5\"",
			},
		},
		{
			name:     "set_query",
			rawQuery: "SET val 5",

			mockCompute: func(computeLayer *mocks.ComputeLayerMock) {
				computeLayer.ParseQueryMock.Times(1).
					Expect("SET val 5").
					Return(entity.MakerCommand{
						CommandType: entity.CommandTypeSet,
						Args:        []string{"val", "5"},
					}.Make(t), nil)
			},
			mockStorage: func(storageLayer *mocks.StorageLayerMock) {
				storageLayer.SetMock.Times(1).
					Expect(minimock.AnyContext, "val", "5").
					Return(nil)
			},

			expectedResp: ExecuteResponse{
				message: "success set",
			},
		},
		{
			name:     "delete_query",
			rawQuery: "DEL val",

			mockCompute: func(computeLayer *mocks.ComputeLayerMock) {
				computeLayer.ParseQueryMock.Times(1).
					Expect("DEL val").
					Return(entity.MakerCommand{
						CommandType: entity.CommandTypeDelete,
						Args:        []string{"val"},
					}.Make(t), nil)
			},
			mockStorage: func(storageLayer *mocks.StorageLayerMock) {
				storageLayer.DeleteMock.Times(1).
					Expect(minimock.AnyContext, "val").
					Return(nil)
			},

			expectedResp: ExecuteResponse{
				message: "success delete",
			},
		},
		{
			name:     "failed_parse_query",
			rawQuery: "DEL val",

			mockCompute: func(computeLayer *mocks.ComputeLayerMock) {
				computeLayer.ParseQueryMock.Times(1).
					Expect("DEL val").
					Return(entity.Command{}, errors.New(errorMessage))
			},
			mockStorage: func(_ *mocks.StorageLayerMock) {},

			expectedResp: ExecuteResponse{fmt.Sprintf("invalid query: %s", errorMessage)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := minimock.NewController(t)
			ctx := context.Background()

			computeMock := mocks.NewComputeLayerMock(ctrl)
			tt.mockCompute(computeMock)
			storageMock := mocks.NewStorageLayerMock(ctrl)
			tt.mockStorage(storageMock)
			loggerMock := logger_mock.NewLoggerAliasMock(ctrl)

			database, err := MakeDatabase(storageMock, computeMock, loggerMock)
			require.NoError(t, err)

			resp, err := database.Execute(ctx, tt.rawQuery)
			if tt.expectedErr != "" {
				require.Error(t, err)
				require.Equal(t, tt.expectedErr, err.Error())
				return
			}
			require.NoError(t, err)

			require.Equal(t, tt.expectedResp, resp)
		})
	}
}
