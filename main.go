// Package main implements Veribi CLI.
package main

import (
	"fmt"

	"github.com/vaclav-dvorak/veribi-cli/cmd/veribi"
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

func main() {
	printWelcome()
	veribi.Execute()
}

func printWelcome() {
	for i := 0; i < len(logoSmall); i++ {
		fmt.Printf("\n%s%s%s", blue, logoSmall[i], reset)
	}
	fmt.Print("\n\n")
}
