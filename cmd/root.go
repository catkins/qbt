package cmd

import (
	"fmt"
	"os"

	"github.com/catkins/qbt/pkg/lua"
	"github.com/spf13/cobra"
)

// Execute runs the qbt CLI application
func Execute() {
	if err := root().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func root() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "qbt",
		Short: "qbt is a query interface for Google Cloud BigTable",
		Run:   queryBigtable,
		Args:  cobra.ExactArgs(1),
	}

	return &rootCmd
}

func queryBigtable(cmd *cobra.Command, args []string) {
	env := lua.NewEnvironment()
	source := args[0]
	err := env.Eval(source)
	if err != nil {
		fmt.Printf("error running script: %v\n", err)
	}
}
