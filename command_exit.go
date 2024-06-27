package main

import (
    "os"
)

func callbackExit(cfg *config, args ...string) error {
    // If the user types "exit" then exit the program and do not 
    // show an error code
    os.Exit(0)
    return nil
}
