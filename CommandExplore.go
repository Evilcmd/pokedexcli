package main

import "fmt"

func exploreCommand(cfg *config, location string) error {

	pokeapiClient := cfg.pokeClient

	ListofPokemons, err := pokeapiClient.ListPokemonsInArea(location)

	if err != nil {
		return err
	}

	fmt.Printf("Pokemons in %v area\n", location)
	for _, pokemon := range ListofPokemons {
		fmt.Println(pokemon)
	}

	return nil
}
