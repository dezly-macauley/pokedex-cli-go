package main

import (
    "github.com/dezly-macauley/pokedex-cli-go/internal/pokeapi"
    "log"
    "fmt"
)

func callbackMap() {

    pokeapiClient := pokeapi.NewClient()

    resp, err := pokeapiClient.ListLocationAreas()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(resp)

}
