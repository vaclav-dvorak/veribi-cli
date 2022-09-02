// Package veribi implements utility routines for manipulating information
// from Veribi platform.
package veribi

import (
	log "github.com/sirupsen/logrus"
)

// GetLatestVersion - returns veribi cli latest version string
func GetLatestVersion() (ver string) {
	rel, err := getLatestRelease()
	if err != nil {
		log.Fatal(err)
	}
	ver = rel.Tag
	return
}
