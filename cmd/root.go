package cmd

import (
	"fmt"
	"os"

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
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello qbt")
		},
	}

	return &rootCmd
}
