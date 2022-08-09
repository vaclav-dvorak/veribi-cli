package veribi

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const host = "https://app.veribi.com/"

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

func scrapeURL(url string) *goquery.Document {
	authCookie := &http.Cookie{
		Name:  "PHPSESSID",
		Value: viper.GetString("key"),
		// Value:  "foo",
		MaxAge: 3600,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(authCookie)
	req.Header.Set("User-Agent", "VeribiCLI - We read content because we don't have API")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if terr := res.Body.Close(); terr != nil {
			log.Fatal(terr)
		}
	}()
	// log.Infof("status: %d", res.StatusCode)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

// ScrapeOffers return list of current offers
func ScrapeOffers(incAuctions bool) (result []Offer, err error) {

	doc := scrapeURL(host + "offers")
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

	doc := scrapeURL(host + off.URL)
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
