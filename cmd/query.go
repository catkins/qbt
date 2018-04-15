package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/catkins/qbt/pkg/bt"
	"github.com/spf13/cobra"
)

func queryCommand(config *flags) cobra.Command {
	return cobra.Command{
		Use:   "query [table] [query]",
		Short: "query data from Google Cloud Bigtable",
		Run:   queryBigtable(config),
		Args:  cobra.ExactArgs(2),
	}
}

func queryBigtable(config *flags) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := bt.NewClient(ctx, config.projectID, config.instance)
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
}
