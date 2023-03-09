// Package veribi implements command line commands that are user inside Veribi CLI.
package veribi

import (
	"fmt"
)

const (
	blue  = "\033[94m"
	reset = "\033[0m"
)

var (
	logoSmall = []string{
		`                 o|    o        |    o`,
		` .    ,,---.,---..|---..   ,---.|    .`,
		`  \  / |---'|    ||   ||---|    |    |`,
		"   `'  `---'`    ``---'`   `---'`---'`",
	}
)

func printLogo() {
	for i := 0; i < len(logoSmall); i++ {
		fmt.Printf("\n%s%s%s", blue, logoSmall[i], reset)
	}
	fmt.Print("\n\n")
}
