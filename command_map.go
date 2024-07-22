package main

import (
	"fmt"
	"log"
)

func commandMap(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	resp, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas: ")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return err
}
