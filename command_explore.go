package main

import (
	"errors"
	"fmt"
)

// NOTE: The input is a variadic string parameter
// In simple terms the user's raw input will be accepted as a slice of strings
func callbackExplore(cfg *config, args ...string) error {

    // TODO: Check this line...
    if len(args) != 1 {
       return errors.New("no location area provided") 
    }

    // The first part of what the user types should be the location area
    locationAreaName := args[0]

    locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
    if err != nil {
        return err
    }
    
    fmt.Printf("Pokemon in %s:\n", locationArea.Name)

    for _, pokemon := range locationArea.PokemonEncounters {
        fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
    }

    return nil

}
