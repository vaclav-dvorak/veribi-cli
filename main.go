// Package main implements Veribi CLI.
package main

import "github.com/vaclav-dvorak/veribi-cli/cmd/veribi"

func main() {
	printWelcome()
	veribi.Execute()
}
