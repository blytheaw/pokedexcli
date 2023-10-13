package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/blytheaw/pokedexcli/internal/pokeapi"
)

type config struct {
	pokedex       map[string]pokeapi.Pokemon
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type command struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func main() {
	fmt.Println("\nWelcome to the Pokedex CLI!")
	fmt.Println("\nType \"help\" for usage instructions or \"exit\" to quit.")

	conf := &config{
		pokedex:       make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeapi.NewClient(30*time.Second, 30*time.Second),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\npokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cmdParts := strings.Split(text, " ")

		cmd, ok := getCommands()[cmdParts[0]]
		if !ok {
			fmt.Println("\nUnknown command. Please try again or type \"help\" for usage instructions.")
		} else {
			err := cmd.callback(conf, cmdParts[1:])

			if err != nil {
				fmt.Println(err)
			}
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
		"explore": {
			name:        "explore",
			description: "Displays Pokemon located at chosen location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View details about one of your caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}
