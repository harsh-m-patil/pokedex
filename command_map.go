package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	resp, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return err
}

func commandMapBack(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	if cfg.prevLocationAreaURL == nil {
		return fmt.Errorf("you are on the first page")
	}
	resp, err := pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return err
}
