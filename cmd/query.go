package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/catkins/qbt/pkg/bt"
	"github.com/catkins/qbt/pkg/lua"
	"github.com/spf13/cobra"
)

func queryBigtable(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	config := loadConfig()

	var client *bt.Client
	var err error

	if config.Emulator != "" {
		client, err = bt.NewClient(ctx, "proj", "instance", bt.WithEmulator(config.Emulator))
	} else {
		client, err = bt.NewClient(ctx, config.Project, config.Instance)
	}

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	luaenv := lua.NewEnvironment()

	table := args[0]
	luaQuery := args[1]

	query := bt.Query{
		Table:     table,
		Range:     bt.AllRows{},
		Predicate: luaenv.RowPredicate(luaQuery),
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
