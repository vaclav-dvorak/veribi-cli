// Package veribi implements utility routines for manipulating information
// from Veribi platform.
package veribi

import (
	"errors"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func refreshKey() (err error) {
	res, err := pingVeribi()
	prevKey := viper.Get("key")
	for _, cookie := range res.Cookies() {
		if cookie.Name == "PHPSESSID" {
			viper.Set("key", cookie.Value)
		}
	}
	if prevKey == viper.Get("key") || viper.Get("key") == "" {
		err = errors.New("cannot get key from Veribi platform")
		return
	}
	return
}

func handle2FA() (err error) {
	prompt := promptui.Prompt{
		Label: "Enter 2FA code",
		Validate: func(input string) error {
			if len(input) != 6 {
				return errors.New("invalid value")
			}
			return nil
		},
	}
	code, err := prompt.Run()
	if err != nil {
		return
	}
	payload := url.Values{}
	payload.Set("code", code)
	res, err := callVeribi("s/2fauth.php", payload)
	defer func() {
		if terr := res.Body.Close(); terr != nil {
			log.Fatal(terr)
		}
	}()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	if doc.Find("title").Text() == "Two Phase Authentication" {
		err = errors.New("incorrect 2FA code")
	}
	return
}

// Login solves whole login process and return error if anything doesn't go as planned
func Login() (err error) {
	if err = refreshKey(); err != nil {
		return
	}

	payload := url.Values{}
	payload.Set("email", viper.GetString("email"))
	payload.Set("pass", viper.GetString("pass"))

	res, err := callVeribi("s/verify.php", payload)
	defer func() {
		if terr := res.Body.Close(); terr != nil {
			log.Fatal(terr)
		}
	}()

	// respDump, err := httputil.DumpResponse(res, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("RESPONSE:\n%s", string(respDump))

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	switch title := doc.Find("title").Text(); title {
	case "Login":
		err = errors.New("incorrect email or password. run veribi init to set them correctly")
		return
	case "Two Phase Authentication":
		err = handle2FA()
	}

	return
}
