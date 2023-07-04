// Package veribi implements command line commands that are user inside Veribi CLI.
package veribi

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vaclav-dvorak/veribi-cli/pkg/veribi"
)

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"l"},
	Short:   "Login into Veribi platform",
	Run: func(cmd *cobra.Command, args []string) {
		if !silent {
			printLogo()
		}
		if viper.GetString("email") == "" || viper.GetString("pass") == "" {
			log.Fatal("email or password is empty, run veribi init")
		}
		if err := veribi.Login(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("🎉 Congratulations you've been logged in. 🎉")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
