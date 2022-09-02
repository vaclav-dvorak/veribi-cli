// Package veribi implements command line commands that are user inside Veribi CLI.
package veribi

import (
	"fmt"
	"sort"

	"github.com/cheynewallace/tabby"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vaclav-dvorak/veribi-cli/pkg/veribi"
)

var sortByThs bool
var incAuctions bool
var offersCmd = &cobra.Command{
	Use:     "offers",
	Aliases: []string{"o"},
	Short:   "List current offers from Veribi platform",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("cmd offers")
		if viper.GetString("key") == "" {
			log.Fatal("run veribi login before running this command")
		}
		offers, err := veribi.ScrapeOffers(incAuctions)
		if err != nil {
			log.Fatal("key is expired or invalid. run veribi login")
		}
		fmt.Print(" Scraping details of offers 0%", len(offers))
		for i := 0; i < len(offers); i++ {
			fmt.Printf("\x1b[3D%02d%%", ((i+1)*100)/len(offers))
			offers[i], err = veribi.ScrapeOffer(offers[i])
			if err != nil {
				log.Fatal("key is expired or invalid. run veribi login")
			}
		}
		fmt.Print("\n")

		if sortByThs {
			sort.Slice(offers, func(i, j int) bool {
				return offers[i].ThsPrice < offers[j].ThsPrice
			})
		}

		t := tabby.New()
		t.AddHeader("ID", "TH price ($/TH)", "Hosting price ($/day)", "Health", "Title")
		for i := 0; i < len(offers); i++ {
			hpf, hp := getHealth(offers[i])
			style := color.New(color.FgGreen).SprintFunc()
			switch { // we will change color based on percent of miners online
			case hpf < 0.1:
				style = color.New(color.FgRed).SprintFunc()
			case hpf < 0.8:
				style = color.New(color.FgYellow).SprintFunc()
			}
			t.AddLine(offers[i].ID, fmt.Sprintf("%6.2f", offers[i].ThsPrice), fmt.Sprintf("%4.2f", offers[i].HostPrice), style(fmt.Sprintf("%3.1d%%", hp)), offers[i].Title)
		}
		t.Print()
	},
}

func init() {
	offersCmd.Flags().BoolVarP(&sortByThs, "ths", "t", false, "Sort by THS/$")
	offersCmd.Flags().BoolVarP(&incAuctions, "add-auctions", "a", false, "Add auctions to the list")
	rootCmd.AddCommand(offersCmd)
}

func getHealth(o veribi.Offer) (hpf float64, hpp int) {
	hpf = (float64(o.Count) - float64(o.NotWorking)) / float64(o.Count)
	hpp = int(hpf * 100)
	return
}
