// Package veribi implements utility routines for manipulating information
// from Veribi platform.
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
	ID         string
	Title      string
	Kind       string
	URL        string
	Count      int
	NotWorking int
	HostPrice  float64
	ThsPrice   float64
}

// ScrapeOffers return list of current offers
func ScrapeOffers(incAuctions bool) (result []Offer, err error) {

	doc, err := scrapeURL("offers")
	if err != nil {
		return
	}
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
const countRegex string = `^Miners: ([0-9]+).*$`
const notWorkingRegex string = `^\(([0-9]+) miners currently not working\)$`

// ScrapeOffer enrich Offer with additional data
func ScrapeOffer(off Offer) (result Offer, err error) {
	doc, err := scrapeURL(off.URL)
	if err != nil {
		return
	}
	if doc.Find("title").Text() == "Login" { //! we have been redirected
		err = errors.New("key expired")
		return
	}

	result = off
	result.ID = strings.Split(off.URL, "=")[1]

	thsRe := regexp.MustCompile(thsPriceRegex)
	hostRe := regexp.MustCompile(hostPriceRegex)
	countRe := regexp.MustCompile(countRegex)
	nwRe := regexp.MustCompile(notWorkingRegex)

	doc.Find(".pb-4 div").Each(func(i int, s *goquery.Selection) {
		match := thsRe.FindStringSubmatch(s.Text())
		if len(match) != 0 {
			result.ThsPrice, err = strconv.ParseFloat(match[1], 64)
		}

		match = countRe.FindStringSubmatch(s.Text())
		if len(match) != 0 {
			result.Count, err = strconv.Atoi(match[1])
		}

		match = hostRe.FindStringSubmatch(s.Text())
		if len(match) != 0 {
			result.HostPrice, err = strconv.ParseFloat(match[1], 64)
		}
	})

	result.NotWorking = 0
	doc.Find(".tg_miners_nw").Each(func(i int, s *goquery.Selection) {
		match := nwRe.FindStringSubmatch(s.Text())
		if len(match) != 0 {
			result.NotWorking, err = strconv.Atoi(match[1])
		}
	})
	return
}
