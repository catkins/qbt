package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type flags struct {
	instance  string
	projectID string
	table     string
}

// Execute runs the qbt CLI application
func Execute() {
	if err := rootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootCommand() *cobra.Command {
	config := flags{}
	rootCmd := cobra.Command{
		Use:   "qbt",
		Short: "QBT is a query cli for Google Cloud BigTable",
		Long: `QBT is a query cli for Google Cloud BigTable.
See https://github.com/catkins/qbt for more information and documentation.`,
	}
	rootCmd.PersistentFlags().StringVarP(&config.instance, "instance", "i", "", "BigTable instance to connect to")
	rootCmd.PersistentFlags().StringVarP(&config.projectID, "project", "p", "", "GCP project to connect to")

	queryCommand := queryCommand(&config)
	rootCmd.AddCommand(&queryCommand)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "interactively initialise configuration for QBT",
		Run:   initConfig,
	})

	return &rootCmd
}

func initConfig(cmd *cobra.Command, args []string) {
	fmt.Println("not yet implemented")
}
