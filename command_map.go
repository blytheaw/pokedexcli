package main

import (
	"fmt"
)

func commandMapForward(config *config, params []string) error {
	locations, err := config.pokeapiClient.ListLocations(config.Next)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapBack(config *config, params []string) error {
	if config.Previous == nil {
		fmt.Println("\nCan't go back any further! Try the 'map' command instead.")
		return nil
	}

	locations, err := config.pokeapiClient.ListLocations(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}
