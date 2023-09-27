package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/blytheaw/pokedexcli/internal/pokecache"
)

func main() {
	fmt.Println("\nWelcome to the Pokedex CLI!")
	fmt.Println("\nType \"help\" for usage instructions or \"exit\" to quit.")
	fmt.Println()
	fmt.Print("pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)

	conf := config{
		Next: "https://pokeapi.co/api/v2/location",
	}
	commands := loadCommands()
	cache := pokecache.NewCache(time.Duration(5) * time.Second)

	for scanner.Scan() {
		text := scanner.Text()

		cmd, ok := commands[text]
		if !ok {
			fmt.Println("\nUnknown command. Please try again or type \"help\" for usage instructions.")
		} else {
			cmd.callback(&conf, &cache)
		}

		fmt.Print("\npokedex > ")
	}
}
