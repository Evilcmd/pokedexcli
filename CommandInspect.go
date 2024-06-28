package main

import "fmt"

func inspectCommand(cfg *config, pokemonName string) error {
	if len(pokemonName) == 0 {
		return fmt.Errorf("pokemon name not entered")
	}

	pokemonInfo, ok := cfg.pokemonsCought[pokemonName]
	if !ok {
		fmt.Printf("you have not caught %v\n", pokemonName)
		return nil
	}

	// fmt.Println(pokemonInfo)
	fmt.Printf("Name: %v\n", pokemonInfo.Name)
	fmt.Printf("Height: %v\n", pokemonInfo.Height)
	fmt.Printf("Weight: %v\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf(" -%v: %v\n", stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ptype := range pokemonInfo.PokemonTypes {
		fmt.Printf(" -%v\n", ptype)
	}

	return nil
}
