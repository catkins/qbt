package cmd

import (
	"github.com/spf13/viper"
)

// Config is the application configuration, provided by CLI flags, ENV vars or configuration file
type Config struct {
	// Project is the GCP project containing the BigTable instance you want to connect to
	Project string
	// Instance is the BigTable instance to connect to
	Instance string

	// Emulator is the address of a BigTable emulator
	// instance and project are ignored when an emulator is provided
	Emulator string

	// Prefix is the row key prefix to predicate upon table scan
	Prefix string
}

func loadConfig() Config {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("qbt")

	config := Config{}
	viper.Unmarshal(&config)
	return config
}
