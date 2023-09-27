package main

import (
	"fmt"
	"os"

	"github.com/blytheaw/pokedexcli/internal/pokeapi"
	"github.com/blytheaw/pokedexcli/internal/pokecache"
)

type config struct {
	Next     string
	Previous *string
}

type command struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
}

func commandHelp(config *config, cache *pokecache.Cache) error {
	commands := loadCommands()

	fmt.Println("\nUsage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandExit(config *config, cache *pokecache.Cache) error {
	os.Exit(0)

	return nil
}

func commandMap(config *config, cache *pokecache.Cache) error {
	locations, err := pokeapi.GetLocations(config.Next, cache)
	if err != nil {
		return err
	}

	config.Next = *locations.Next
	if locations.Previous != nil {
		config.Previous = locations.Previous
	} else {
		config.Previous = nil
	}

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapBack(config *config, cache *pokecache.Cache) error {
	if config.Previous == nil {
		fmt.Println("\nCan't go back any further! Try the 'map' command instead.")
		return nil
	}

	locations, err := pokeapi.GetLocations(*config.Previous, cache)
	if err != nil {
		return err
	}

	config.Next = *locations.Next
	if locations.Previous != nil {
		config.Previous = locations.Previous
	} else {
		config.Previous = nil
	}

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func loadCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex CLI",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations on the map",
			callback:    commandMapBack,
		},
	}
}
