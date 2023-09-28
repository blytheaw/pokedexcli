package main

import (
	"fmt"
)

func commandHelp(config *config) error {
	fmt.Println("\nUsage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
