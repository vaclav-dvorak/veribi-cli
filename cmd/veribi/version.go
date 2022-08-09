package veribi

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Gets CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Veribi CLI Version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
