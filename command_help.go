package main

import (
    "fmt"
)

func callbackHelp() error {
    fmt.Println("Welcome to the pokedex help menu")
    fmt.Println("===============================")
    fmt.Println("Here are your available commands:")

    for _, cmd := range getCommands() {
        fmt.Printf(" - %s: %s\n", cmd.name, cmd.description )
    }

    fmt.Println("")
    return nil

}
