package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NOTE: The reason this is a pointer is because you want the entire app
// to have access to a reference of this config instance.
// This will prevent wasted memory from accidentally copying the data 
// structure

func startRepl(cfg *config) {

    // infinite for loop
    for {

        // create a scanner to capture the user input from the terminal
        scanner := bufio.NewScanner(os.Stdin)

        // prompt the user to seach for something
        fmt.Print("Type your search query >> ")

        // activate the scanner
        scanner.Scan()

        // save the user input into a variable
        text := scanner.Text()

        // Clean the user input before printing it back
        cleaned := cleanInput(text)

        // NOTE: What to do if user presses enter without searching for 
        // anything

        if len(cleaned) == 0 {
            // continue looping
            continue            
        }

        // This first word in the user's input is a command
        commandName := cleaned[0]

        // a variable to store a map of the available commands
        availableCommands := getCommands()

        command, ok := availableCommands[commandName]

        // NOTE: Error handling for a situation where the user enters a 
        // command that does not exist in the map of available Commands 
        if !ok {
            fmt.Println("Invalid command")
            continue
        }
        err := command.callback(cfg)

        if err != nil {
            fmt.Println(err)
        }
        

    }

}

type cliCommand struct {

    // Every command will have a name, description, and a function that
    // will perform its task
    name string
    description string
    callback func(*config) error
}

// This function will return a map where each key in the map 
// (which is a string) is paired with a variable that is of the type 
// cliCommand struct
func getCommands() map[string]cliCommand {

    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Open the help menu",
            callback: callbackHelp, 
        },
        "map": {
            name: "map",
            description: "Shows the next page of locations",
            callback: callbackMap, 
        },
        "mapb": {
            name: "mapb",
            description: "Shows the previous page of locations",
            callback: callbackMapb, 
        },
        "exit": {
            name: "exit",
            description: "Close the Pokedex",
            callback: callbackExit,
        },

    }

}

// This will make a few changes to the raw input that the user types
// to make it ready for processing
func cleanInput(str string) []string {

    // This will convert all the letters of the output to lowercase
    lowered := strings.ToLower(str)

    // This will return a []string
    // Basically a slice of strings where every word is a seperate element
    words :=  strings.Fields(lowered)

    return words

}
