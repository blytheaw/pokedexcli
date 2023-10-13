package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, params []string) error {
	if len(params) != 1 {
		return errors.New("\nExpected format: 'inspect <pokemon>'")
	}

	pokemonName := params[0]
	pokemon, ok := config.pokedex[pokemonName]

	if !ok {
		fmt.Println("\nYou have not caught that Pokemon! Try catching it first.")
		return nil
	}

	fmt.Printf("\nName: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	return nil
}
