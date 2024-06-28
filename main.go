package main

import "github.com/Evilcmd/pokedexcli/internal/pokeapi"

// holds stateful information
type config struct {
	pokeClient              pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
	pokemonsCought          map[string]pokeapi.PokemonInfoDefn
}

func main() {

	cfg := config{
		pokeClient:     pokeapi.NewClient(),
		pokemonsCought: make(map[string]pokeapi.PokemonInfoDefn),
	}

	startRepl(&cfg)
}
