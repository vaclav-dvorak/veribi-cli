package veribi

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	Short:   "Update Veribi cli to newest version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("cmd update")
		fmt.Println("TODO - implement this")
		// curl -v https://api.github.com/repos/vaclav-dvorak/veribi-cli/releases/latest
		// jsonunmarshal "tag_name"
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
