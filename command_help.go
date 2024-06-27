package main

import (
    "fmt"
)

func callbackHelp(cfg *config) error {
    fmt.Println("")
    fmt.Println("--- Welcome to the Pokedex help menu ---")
    fmt.Println("==========================================")
    fmt.Println("Here are your available commands:")

    for _, cmd := range getCommands() {
        fmt.Printf(" - %s: %s\n", cmd.name, cmd.description )
    }
    fmt.Println("==========================================")
    fmt.Println("")
    return nil

}
