// Package veribi implements command line commands that are user inside Veribi CLI.
package veribi

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
	"github.com/vaclav-dvorak/veribi-cli/pkg/veribi"
)

type versionOutput struct {
	VersionString string `json:"version"`
	Outdated      bool   `json:"outdated"`
}

var (
	outputJSON bool
	version    = "v0.0.1+sha"

	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Gets CLI version",
		Run: func(_ *cobra.Command, _ []string) {
			if !silent && !outputJSON {
				printLogo()
			}

			latest := veribi.GetLatestVersion()

			outdated := semver.New(version[1:]).LessThan(*semver.New(latest[1:])) //? we are striping "v" from our tags as its not technically part of semver
			if outputJSON {
				o, err := json.Marshal(versionOutput{VersionString: version, Outdated: outdated})
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s", o)
				return
			}
			fmt.Printf("Veribi CLI Version: %s\n", version)
			if outdated {
				fmt.Printf("\nYour version of Veribi CLI is out of date! The latest version\n"+
					"is %s. You can update by downloading from https://github.com/%s/releases\n", latest[1:], veribi.VeribiRepo)
			}
		},
	}
)

func init() {
	versionCmd.Flags().BoolVarP(&outputJSON, "json", "j", false, "Get output as JSON object")
	rootCmd.AddCommand(versionCmd)
}
