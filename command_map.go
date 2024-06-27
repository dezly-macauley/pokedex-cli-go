package main

import (
    "github.com/dezly-macauley/pokedex-cli-go/internal/pokeapi"
    "log"
    "fmt"
)

func callbackMap() error {

    pokeapiClient := pokeapi.NewClient()

    resp, err := pokeapiClient.ListLocationAreas()
    if err != nil {
        log.Fatal(err)
    }

    for _, area := range resp.Results {
        fmt.Printf(" - %s\n", area.Name)
    }

    return nil

}
