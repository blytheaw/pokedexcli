package main

type command struct {
    name string
    description string
    callback func() error
}

func commandHelp() error {
    return nil
}

func commandExit() error {
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
