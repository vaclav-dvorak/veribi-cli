// Package veribi implements utility routines for manipulating information
// from Veribi platform.
package veribi

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const host = "https://app.veribi.com/"

func scrapeURL(uri string) (doc *goquery.Document, err error) {
	res, err := callAuthVeribi(uri)
	if err != nil {
		return
	}

	defer func() {
		if terr := res.Body.Close(); terr != nil {
			log.Fatal(terr)
		}
	}()
	// log.Infof("status: %d", res.StatusCode)
	doc, err = goquery.NewDocumentFromReader(res.Body)
	return
}

func pingVeribi() (res *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "VeribiCLI - We read content because we don't have API")

	res, err = http.DefaultClient.Do(req)

	return
}

func authCookie() *http.Cookie {
	return &http.Cookie{
		Name:   "PHPSESSID",
		Value:  viper.GetString("key"),
		MaxAge: 3600,
	}
}

func callVeribi(uri string, payload url.Values) (res *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, host+uri, strings.NewReader(payload.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "VeribiCLI - We read content because we don't have API")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(authCookie())

	// reqDump, err := httputil.DumpRequestOut(req, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("REQUEST:\n%s", string(reqDump))

	res, err = http.DefaultClient.Do(req)

	return
}

func callAuthVeribi(uri string) (res *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, host+uri, nil)
	if err != nil {
		return
	}
	req.AddCookie(authCookie())
	req.Header.Set("User-Agent", "VeribiCLI - We read content because we don't have API")

	res, err = http.DefaultClient.Do(req)

	return
}
