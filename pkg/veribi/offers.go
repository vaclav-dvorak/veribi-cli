package veribi

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Offer is a struct holding information about current offer
type Offer struct {
	ID        string
	Title     string
	Kind      string
	URL       string
	Count     int
	HostPrice float64
	ThsPrice  float64
}

// ScrapeOffers return list of current offers
func ScrapeOffers(incAuctions bool) (result []Offer, err error) {

	doc := scrapeURL("offers")
	if doc.Find("title").Text() == "Login" { //! we have been redirected
		err = errors.New("key expired")
		return
	}
	result = make([]Offer, 0)

	// lv-openoff
	// lv-auctions
	// lv-futureoff

	doc.Find(".itmboxin .of_title").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")
		kind := "offer"
		if strings.Contains(href, "auction") {
			kind = "auction"
		}
		if incAuctions || kind != "auction" {
			result = append(result, Offer{Title: s.Find("a").Text(), Kind: kind, URL: href})
		}
	})

	return
}

const thsPriceRegex string = `^Price per miner: \$[0-9,\.]* \(\$([0-9,\.]+)\/TH\)$`
const hostPriceRegex string = `^Hosting per day for one miner: \$([0-9\.]+)$`
const countRegex string = `^Miners: ([0-9]+) .*$`

// ScrapeOffer enrich Offer with additional data
func ScrapeOffer(off Offer) (result Offer, err error) {

	doc := scrapeURL(off.URL)
	if doc.Find("title").Text() == "Login" { //! we have been redirected
		err = errors.New("key expired")
		return
	}

	result = off
	result.ID = strings.Split(off.URL, "=")[1]

	thsRe := regexp.MustCompile(thsPriceRegex)
	hostRe := regexp.MustCompile(hostPriceRegex)
	countRe := regexp.MustCompile(countRegex)

	doc.Find(".pb-4 div").Each(func(i int, s *goquery.Selection) {
		match1 := thsRe.FindStringSubmatch(s.Text())
		if len(match1) != 0 {
			result.ThsPrice, err = strconv.ParseFloat(match1[1], 64)
		}

		match2 := countRe.FindStringSubmatch(s.Text())
		if len(match2) != 0 {
			result.Count, err = strconv.Atoi(match2[1])
		}

		match3 := hostRe.FindStringSubmatch(s.Text())
		if len(match3) != 0 {
			result.HostPrice, err = strconv.ParseFloat(match3[1], 64)
		}
	})
	return
}
