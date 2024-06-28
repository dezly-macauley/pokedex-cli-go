package main

import (
	// "crypto/rand"
	"errors"
	"fmt"
)

// NOTE: The input is a variadic string parameter
// In simple terms the user's raw input will be accepted as a slice of strings
func callbackInspect(cfg *config, args ...string) error {

    // TODO: Check this line...
    if len(args) != 1 {
       return errors.New("no pokemon name provided") 
    }

    // The first part of what the user types should be the location area
    pokemonName := args[0]

    pokemon, ok := cfg.caughtPokemon[pokemonName]
    if !ok {
        return errors.New("You have not caught this pokemon yet")
    }

    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %v\n", pokemon.Height)
    fmt.Printf("Weight: %v\n", pokemon.Weight)
    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf(" - %s: %v", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Println("Types:")
    for _, typ := range pokemon.Types {
        fmt.Printf(" - %s\n", typ.Type.Name)
    }

    return nil

}
