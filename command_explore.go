package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, params []string) error {
	if len(params) != 1 {
		return errors.New("Expected format: 'explore <location>'")
	}

	locationName := params[0]

	fmt.Println("Exploring " + locationName + "...")
	location, err := config.pokeapiClient.GetLocationByName(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, p := range location.PokemonEncounters {
		fmt.Println(" - " + p.Pokemon.Name)
	}

	return nil
}
