package main

import (
	"time"
	"github.com/dezly-macauley/pokedex-cli-go/internal/pokeapi"
)

// This will hold all of the stateful information for the command callback
// function
type config struct {
    pokeapiClient           pokeapi.Client

    // The reason these are pointers to a string,
    // is to account for the possibility that these could be nil
    // E.g. Situations where there is no next page, or previous page
    nextLocationAreaURL     *string
    previousLocationAreaURL *string
    caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {

    cfg := config {
        pokeapiClient: pokeapi.NewClient(time.Hour),
        caughtPokemon: make(map[string]pokeapi.Pokemon),
    }

    // Prompt the user for their input, evaluate it, and print it
    startRepl(&cfg)

}
