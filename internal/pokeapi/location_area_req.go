package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// NOTE: Methods that new instances created from the "Client" struct can use

// You could rename c to instanceOfClient if you wanted to be more verbose
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error){
    
    endpoint := "/location-area"
    fullURL := baseURL + endpoint

    if pageURL != nil {
        fullURL = *pageURL
    }
    
    dat, ok := c.cache.Get(fullURL)
    if ok {

        // fmt.Println("cache hit!")
        locationAreasResp := LocationAreasResp{}

        // json.unmarshal takes in the data and a pointer to the struct that you
        // want to unmarshal the data into
        err := json.Unmarshal(dat, &locationAreasResp)
        if err != nil {
            return LocationAreasResp{}, err
        }
        
        return locationAreasResp, nil

    }
    // fmt.Println("cache miss!")

    // NOTE: Making a new HTTP request to this endpoint

    // The last arguement is nil, because you are not sending a body
    // I.e. this is a GET request, so data is only being recieved into the 
    // program

    // req is short for "request"
    req, err := http.NewRequest("GET", fullURL, nil)
    
    // If there was a problem making the request then return an empty struct,
    // and the error
    if err != nil {
        return LocationAreasResp{}, err
    }

    // Now that there is a valid request, execute the request
    resp, err := c.httpClient.Do(req)

    // If there is an err executing the the request 
    if err != nil {
        return LocationAreasResp{}, err
    }

    // NOTE: Remember to close the response object when you are done with it.
    defer resp.Body.Close()

    if resp.StatusCode > 399 {
        return LocationAreasResp{}, fmt.Errorf("bad status code: %v",
            resp.StatusCode)
    }

    // NOTE: Next read the data from the response body

    // This will return data, which will be a slice of bytes, and an error
    dat, err = io.ReadAll(resp.Body)
    if err != nil {
        return LocationAreasResp{}, err
    }

    // NOTE: Now that JSON data which is recieved as a slice of bytes needs
    // to be unmarshalled into a struct

    locationAreasResp := LocationAreasResp{}

    // json.unmarshal takes in the data and a pointer to the struct that you
    // want to unmarshal the data into
    err = json.Unmarshal(dat, &locationAreasResp)
    if err != nil {
        return LocationAreasResp{}, err
    }
    
    c.cache.Add(fullURL, dat)

    return locationAreasResp, nil

}
