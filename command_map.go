package main

import (
	"fmt"
	"log"

	"github.com/harsh-m-patil/pokedex/internal/pokeapi"
)

func commandMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas: ")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	return err
}
