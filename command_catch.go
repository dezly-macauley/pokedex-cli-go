package main

import (
	// "crypto/rand"
	"errors"
	"fmt"
    "math/rand"
)

// NOTE: The input is a variadic string parameter
// In simple terms the user's raw input will be accepted as a slice of strings
func callbackCatch(cfg *config, args ...string) error {

    // TODO: Check this line...
    if len(args) != 1 {
       return errors.New("no pokemon name provided") 
    }

    // The first part of what the user types should be the location area
    pokemonName := args[0]

    pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
    if err != nil {
        return err
    }

    // This will generate a random number between 0 and base experience

    const threshold = 50
    randNum := rand.Intn(pokemon.BaseExperience)
    
    if randNum > threshold {
        return fmt.Errorf("Failed to catch %s", pokemonName)
    }

    cfg.caughtPokemon[pokemonName] = pokemon
    fmt.Printf("%s was caught!\n", pokemonName)
    return nil

}
