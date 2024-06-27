package main

// NOTE: To run this file

// To test the main package
// `go test`
// `go test ./...` This will test everything recursively  


import (
    "testing"
)

func TestCleanInput(t *testing.T) {

    cases := []struct {
        input string
        expected []string
    } {

        {
            input: "hello world",
            expected: []string {
                "hello",
                "world",
            },
        },
        
        {
            input: "HELLO WORLD",
            expected: []string {
                "hello",
                "world",
            },
        },

    }


    // cs is short for case
    // _, is because range works with two values
    // index, value := range thing-that-you-want-to-loop-over
    // So _, cs means that you don't care about the index, only the value
    for _, cs := range cases {

        actual := cleanInput(cs.input)
        if len(actual) != len(cs.expected) {
            t.Errorf("The lengths are not equal: %v vs %v",
                len(actual),
                len(cs.expected),
            )
            continue
        }
        for i := range actual {
            actualWord := actual[i]
            expectedWord := cs.expected[i]
            if actualWord != expectedWord {
                t.Errorf("%v does not equal %v", actualWord, expectedWord)
            }
        }

    }

}
