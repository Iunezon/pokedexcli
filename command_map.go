package main

import (
	"fmt"
	"log"

	"github.com/Iunezon/pokedexcli/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" – %s\n", area)
	}
	return nil
}
