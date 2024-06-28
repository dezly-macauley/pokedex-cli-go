package main

import (
	"fmt"
)

// NOTE: The input is a variadic string parameter
// In simple terms the user's raw input will be accepted as a slice of strings
func callbackPokedex(cfg *config, args ...string) error {

    fmt.Println("")
    fmt.Println("Pokemon in Pokedex:")
    fmt.Println("==============================")
    for _, pokemon := range cfg.caughtPokemon {
        fmt.Printf(" - %s\n", pokemon.Name)
    }
    fmt.Println("==============================")
    fmt.Println("")

    return nil

}
