package main

import (
	// "github.com/dezly-macauley/pokedex-cli-go/internal/pokeapi"
	"errors"
	"fmt"
	"log"
)

// NOTE: Each time this function is called it will load the next page of
// locations
func callbackMap(cfg *config) error {

    resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
    if err != nil {
        log.Fatal(err)
    }

    for _, area := range resp.Results {
        fmt.Printf(" - %s\n", area.Name)
    }

    cfg.nextLocationAreaURL = resp.Next
    cfg.previousLocationAreaURL = resp.Previous

    return nil

}

// NOTE: This one will go back to the previous page of Pokemon locations
func callbackMapb(cfg *config) error {

    if cfg.previousLocationAreaURL == nil {
        return errors.New("No previous page. You are already on the first page")
    }

    resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
    if err != nil {
        log.Fatal(err)
    }

    for _, area := range resp.Results {
        fmt.Printf(" - %s\n", area.Name)
    }

    cfg.nextLocationAreaURL = resp.Next
    cfg.previousLocationAreaURL = resp.Previous

    return nil

}
