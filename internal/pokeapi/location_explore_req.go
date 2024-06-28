package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonsInArea(location string) ([]string, error) {

	endpoint := "/location-area/"
	fullurl := BaseUrl + endpoint + location

	body, ok := c.cache.Get(fullurl)
	if ok {
		fmt.Println("Cache Hit")

		LocationExploreRes := LocationExploreDefn{}
		err := json.Unmarshal(body, &LocationExploreRes)
		if err != nil {
			return []string{}, err
		}

		ListOfPokemons := make([]string, 0, len(LocationExploreRes.PokemonEncounters))

		for _, pokemonEnc := range LocationExploreRes.PokemonEncounters {
			ListOfPokemons = append(ListOfPokemons, pokemonEnc.Pokemon.Name)
		}

		return ListOfPokemons, nil

	}

	fmt.Println("Cache Miss")

	req, err := http.NewRequest("GET", fullurl, nil)

	if err != nil {
		return []string{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []string{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println(fullurl)
		return []string{}, fmt.Errorf("error: returned with staus code %v", res.StatusCode)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}

	LocationExploreRes := LocationExploreDefn{}
	err = json.Unmarshal(body, &LocationExploreRes)
	if err != nil {
		return []string{}, err
	}

	ListOfPokemons := make([]string, 0, len(LocationExploreRes.PokemonEncounters))

	for _, pokemonEnc := range LocationExploreRes.PokemonEncounters {
		ListOfPokemons = append(ListOfPokemons, pokemonEnc.Pokemon.Name)
	}

	c.cache.Add(fullurl, body)

	return ListOfPokemons, nil
}
