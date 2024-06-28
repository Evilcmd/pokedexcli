package main

import (
	"fmt"
)

func mapCommand(cfg *config) error {
	pokeapiClient := cfg.pokeClient

	res, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)

	if err != nil {
		return err
	}

	cfg.nextLocationAreaUrl = res.Next
	cfg.previousLocationAreaUrl = res.Previous

	fmt.Println("Location Areas")
	for _, v := range res.Results {
		fmt.Printf(" - %v\n", v.Name)
	}

	return nil
}

func mapbCommand(cfg *config) error {
	pokeapiClient := cfg.pokeClient

	if cfg.previousLocationAreaUrl == nil {
		return fmt.Errorf("cannot go back, you are on first page")
	}

	res, err := pokeapiClient.ListLocationAreas(cfg.previousLocationAreaUrl)

	if err != nil {
		return err
	}

	cfg.nextLocationAreaUrl = res.Next
	cfg.previousLocationAreaUrl = res.Previous

	fmt.Println("Location Areas")
	for _, v := range res.Results {
		fmt.Printf(" - %v\n", v.Name)
	}

	return nil
}
