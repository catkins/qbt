package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/catkins/qbt/pkg/bt"
	"github.com/spf13/cobra"
)

func queryBigtable(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	config := loadConfig()
	client, err := bt.NewClient(ctx, config.Project, config.Instance)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	table := args[0]

	query := bt.Query{
		Table: table,
		Range: bt.AllRows{},
		Predicate: func(row bt.Row) (bool, error) {
			return true, nil
		},
	}

	err = client.ReadRowsFiltered(ctx, query, func(row bt.Row) {
		rowBytes, err := json.Marshal(row)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Printf("%s\n", rowBytes)
	})

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
