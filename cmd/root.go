package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Execute runs the qbt CLI application
func Execute() {
	if err := rootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "qbt",
		Short: "QBT is a query cli for Google Cloud BigTable",
		Long: `QBT is a query cli for Google Cloud BigTable.

See https://github.com/catkins/qbt for more information and documentation.`,
	}
	// global flags
	rootCmd.PersistentFlags().StringP("instance", "i", "", "BigTable instance to connect to")
	viper.BindPFlag("instance", rootCmd.PersistentFlags().Lookup("instance"))
	rootCmd.PersistentFlags().StringP("project", "p", "", "GCP project to connect to")
	viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))

	rootCmd.AddCommand(&cobra.Command{
		Use:   "query [table] [query]",
		Short: "query data from Google Cloud Bigtable",
		Run:   queryBigtable,
		Args:  cobra.ExactArgs(2),
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "interactively initialise configuration for QBT",
		Run:   initConfig,
	})

	return &rootCmd
}
