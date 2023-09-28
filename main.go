package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/blytheaw/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type command struct {
	name        string
	description string
	callback    func(*config) error
}

func main() {
	fmt.Println("\nWelcome to the Pokedex CLI!")
	fmt.Println("\nType \"help\" for usage instructions or \"exit\" to quit.")

	conf := &config{
		pokeapiClient: pokeapi.NewClient(30*time.Second, 30*time.Second),
		//Next:   "https://pokeapi.co/api/v2/location?offset=0&limit=20",
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\npokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cmd, ok := getCommands()[text]
		if !ok {
			fmt.Println("\nUnknown command. Please try again or type \"help\" for usage instructions.")
		} else {
			cmd.callback(conf)
		}
	}
}

func getCommands() map[string]command {
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
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations on the map",
			callback:    commandMapBack,
		},
	}
}
