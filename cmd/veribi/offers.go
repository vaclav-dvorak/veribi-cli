package veribi

import (
	"fmt"
	"sort"

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
		fmt.Printf(" Scraping details of offers 00/%02d\x1b[3D", len(offers))
		for i := 0; i < len(offers); i++ {
			fmt.Printf("\x1b[2D%02d", i+1)
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

		fmt.Printf("| ID  | Price of TH | Hosting price | Title |\n")
		for i := 0; i < len(offers); i++ {
			fmt.Printf("| #%s | %6.2f $/TH | %7.2f $/day | %s\n", offers[i].ID, offers[i].ThsPrice, offers[i].HostPrice, offers[i].Title)
		}

		// fmt.Println("TODO - implement this")
	},
}

func init() {
	offersCmd.Flags().BoolVarP(&sortByThs, "ths", "t", false, "Sort by THS/$")
	offersCmd.Flags().BoolVarP(&incAuctions, "add-auctions", "a", false, "Add auctions to the list")
	rootCmd.AddCommand(offersCmd)
}
