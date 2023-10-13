package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *config, params []string) error {
	if len(params) > 0 {
		return errors.New("Expected format: 'pokedex'")
	}

	fmt.Println("\nYour Pokedex:")

	for n := range config.pokedex {
		fmt.Printf("- %s\n", n)
	}

	return nil
}
