package veribi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// VeribiRepo contain information about our repository on github
const VeribiRepo = "vaclav-dvorak/veribi-cli"
const githubHost = "https://api.github.com/repos/" + VeribiRepo
const timeoutSec = 5

// Release hold parsed information about github release
type Release struct {
	Tag string `json:"tag_name"`
}

func getLatestRelease() (rel Release, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*timeoutSec))
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, githubHost+"/releases/latest", nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	body, _ := io.ReadAll(res.Body)
	defer func() {
		_ = res.Body.Close()
	}()
	err = json.Unmarshal(body, &rel)
	return
}
