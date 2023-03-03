package main

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

func printWelcome() {
	for i := 0; i < len(logoSmall); i++ {
		fmt.Printf("%s%s%s\n", blue, logoSmall[i], reset)
	}
	fmt.Println("")
}
