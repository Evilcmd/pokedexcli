package main

import "fmt"

func pokedexCommand(cfg *config, _ string) error {

	fmt.Println("Your Pokedex:")

	for k := range cfg.pokemonsCought {
		fmt.Printf(" - %v\n", k)
	}

	return nil
}
