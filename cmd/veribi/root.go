// Package veribi implements command line commands that are user inside Veribi CLI.
package veribi

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "veribi",
	Short: "veribi - a simple CLI to interact with Veribi platform",
	Long: `veribi-cli is a simple http envelope to interact with Veribi platform

User can interact with Veribi directly from terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute - initialize and run cobra command
func Execute() {
	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Whoops. There was an error while executing your CLI '%s'", err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// stub init
	configHome, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configName := "config"
	configType := "yaml"
	configPath := filepath.Join(configHome, ".veribi")
	configFile := filepath.Join(configPath, configName+"."+configType)

	if err = os.MkdirAll(configPath, 0755); err != nil {
		log.Fatal(err)
	}
	if _, err = os.OpenFile(configFile, os.O_CREATE, 0644); err != nil {
		log.Fatal(err)
	}

	viper.SetEnvPrefix("veribi") // will be uppercase automatically

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	if err = viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	viper.AutomaticEnv() // read in environment variables that match
}
