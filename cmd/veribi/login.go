package veribi

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"l"},
	Short:   "Login into Veribi platform",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("cmd login")
		if viper.GetString("email") == "" || viper.GetString("pass") == "" {
			log.Fatal("run veribi init before running this command")
		}
		fmt.Println("TODO - implement this")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
