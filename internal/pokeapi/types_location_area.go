package pokeapi

// NOTE: This package will be reponsible for contacting the pokeapi
// over the internet

// Some online tools used to make the JSON easier to read
// 1. JSONLint 
// 2. JSON-to-Go

// NOTE: Do not just blinding copy the output from JSON-to-Go!
// This is great for getting a basic structure down, 
// but the variable types of the fields may not be the most appropriate.
// So you always want to have a look at JSONLint to see what are the correct
// or possible return types

type LocationAreasResp struct {
    Count int `json:"count"`

    // "Next" will give you the next 20 Pokemon locations.
    // The JSON showed that this field could sometimes have the value of null
    // I.e. You have reached the end of the list and there are no more 
    // locations to view
    // In Go the best way to represent something like this is a pointer to 
    // a string. When the value is null, the Field next will be nill
    // Otherwise it will be a pointer that points to a valid string value
    // Same thing goes for "Previous"
    Next *string `json:"next"`
    Previous *string `json:"previous"`

    Results []struct {
        Name string `json:"name"`
        URL string `json:"url"`
    } `json:"results"`

}
