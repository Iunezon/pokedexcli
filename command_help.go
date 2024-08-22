package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to Pokédex help menu!")
	fmt.Println("Here are your available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" — %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}