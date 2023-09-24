package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println("\nWelcome to the Pokedex CLI!")
    fmt.Println("\nType \"help\" for usage instructions or \"exit\" to quit.")
    fmt.Println()
    fmt.Print("pokedex > ")
    scanner := bufio.NewScanner(os.Stdin)

    commands := loadCommands()

    for scanner.Scan() {
        text := scanner.Text()

        cmd, ok := commands[text]
        if !ok {
            fmt.Println("\nUnknown command. Please try again or type \"help\" for usage instructions.")
        } else {
            cmd.callback()
        }

        fmt.Print("\npokedex > ") 
    }
}
