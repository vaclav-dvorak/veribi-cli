package veribi

import (
	"fmt"
)

var (
	version, sha, date = "v0.0.0", "sha", "now"
)

// GetVersion - outputs veribi version string
func GetVersion() (ver string) {
	ver = fmt.Sprintf("Veribi CLI Version: %s (%s) @ %s", version, sha, date)
	return
}
