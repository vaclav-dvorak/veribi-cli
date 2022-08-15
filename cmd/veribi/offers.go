package veribi

import (
	"fmt"
	"sort"

	"github.com/cheynewallace/tabby"
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
		t.AddHeader("ID", "TH price ($/TH)", "Hosting price ($/day)", "Title")
		for i := 0; i < len(offers); i++ {
			t.AddLine(offers[i].ID, fmt.Sprintf("%6.2f", offers[i].ThsPrice), fmt.Sprintf("%4.2f", offers[i].HostPrice), offers[i].Title)
		}
		t.Print()

		// fmt.Println("TODO - implement this")
	},
}

func init() {
	offersCmd.Flags().BoolVarP(&sortByThs, "ths", "t", false, "Sort by THS/$")
	offersCmd.Flags().BoolVarP(&incAuctions, "add-auctions", "a", false, "Add auctions to the list")
	rootCmd.AddCommand(offersCmd)
}
