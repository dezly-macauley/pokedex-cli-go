package pokeapi

import (
	"net/http"
	"time"

	"github.com/dezly-macauley/pokedex-cli-go/internal/pokecache"
)

//=============================================================================

// All of the endpoints in this project will implement this baseURL
const baseURL = "https://pokeapi.co/api/v2"

// NOTE: In Go you need an HTTP client in order to make http requests

type Client struct {
    cache pokecache.Cache
    httpClient http.Client
}

// This is a function that makes it easier to create a new instance of this
// struct with the fields intialized.
//The equivalent of this in a language like TypeScript would 
// be a constructor function. 

func NewClient(cacheInterval time.Duration) Client {

    return Client {
        cache: pokecache.NewCache(cacheInterval),
        httpClient: http.Client {
            // After a minute a given http request will terminated.
            // This prevent the application from grinding to a halt because 
            // something extenal is taking forever to return 
            Timeout: time.Minute,
        }, 

    }

}

