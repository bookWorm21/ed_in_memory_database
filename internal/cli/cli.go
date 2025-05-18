package cli

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"

	"ed_in_memory_database/internal/service/database"
)

// deps -
type (
	databaseService interface {
		Execute(ctx context.Context, rawQuery string) (database.ExecuteResponse, error)
	}
)

type Cli struct {
	databaseService databaseService
}

func Make(databaseService databaseService) (Cli, error) {
	if databaseService == nil {
		return Cli{}, errors.New("database service is nil")
	}
	return Cli{
		databaseService: databaseService,
	}, nil
}

func (c Cli) Run(ctx context.Context) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		request, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		response, err := c.databaseService.Execute(ctx, request)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(response)
	}
}
