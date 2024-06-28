package main

import (
	"fmt"
	"math/rand"
)

func catchCommand(cfg *config, pokemonName string) error {

	if len(pokemonName) == 0 {
		return fmt.Errorf("pokemon name not entered")
	}

	_, ok := cfg.pokemonsCought[pokemonName]
	if ok {
		fmt.Println("Pokemon Already cought")
		return nil
	}

	pokeapiClient := cfg.pokeClient

	pokemonInfo, err := pokeapiClient.GetPokemonInfo(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	const threshold = 50
	randNum := rand.Intn(pokemonInfo.BaseExperience)
	// comment the below line after debug
	// fmt.Println(pokemonInfo.BaseExperience, randNum, threshold)
	if randNum < threshold {
		fmt.Printf("%v was cought\n", pokemonName)
		cfg.pokemonsCought[pokemonName] = pokemonInfo
	} else {
		fmt.Printf("%v escaped\n", pokemonName)
	}

	return nil
}
