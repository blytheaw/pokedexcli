package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Welcome to the Pokedex CLI!")
    fmt.Println("Type \"help\" for usage instructions and \"exit\" to quit.")
    fmt.Println()
    fmt.Print("pokedex > ")
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        text := scanner.Text()
        fmt.Println(text)
        fmt.Print("pokedex > ") 
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
