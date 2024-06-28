package main

import "github.com/Evilcmd/pokedexcli/internal/pokeapi"

// holds stateful information
type config struct {
	pokeClient              pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
}

func main() {

	cfg := config{
		pokeClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
