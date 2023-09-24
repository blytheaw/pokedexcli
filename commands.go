package main

import (
	"fmt"
	"os"
)

type command struct {
    name string
    description string
    callback func() error
}

func commandHelp() error {
    commands := loadCommands()

    fmt.Println("\nUsage:")
    fmt.Println()
    for _, cmd := range commands {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }

    return nil
}

func commandExit() error {
    os.Exit(0)

    return nil
}

func loadCommands() map[string]command {
    return map[string]command {
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the Pokedex CLI",
            callback: commandExit,
        },
    }
}
