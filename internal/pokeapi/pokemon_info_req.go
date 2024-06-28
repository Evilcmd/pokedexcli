package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName string) (PokemonInfoDefn, error) {

	endpoint := "/pokemon/"
	fullurl := BaseUrl + endpoint + pokemonName

	body, ok := c.cache.Get(fullurl)
	if ok {

		// fmt.Println("Cache Hit")

		pokemonApiRes := PokemonInfoAPIDefn{}
		err := json.Unmarshal(body, &pokemonApiRes)
		if err != nil {
			return PokemonInfoDefn{}, err
		}

		numberOfPokemonPastTypes := 0
		for _, past_types := range pokemonApiRes.PastTypes {
			numberOfPokemonPastTypes += len(past_types.Types)
		}

		pokemonInfo := PokemonInfoDefn{
			Name:           pokemonApiRes.Name,
			BaseExperience: pokemonApiRes.BaseExperience,
			Height:         pokemonApiRes.Height,
			Weight:         pokemonApiRes.Weight,
			Stats: make([]struct {
				Name     string
				BaseStat int
			}, 0, len(pokemonApiRes.Stats)),
			PokemonTypes: make([]string, 0, len(pokemonApiRes.Types)+numberOfPokemonPastTypes),
		}

		for _, stat := range pokemonApiRes.Stats {
			pokemonInfo.Stats = append(pokemonInfo.Stats, struct {
				Name     string
				BaseStat int
			}{stat.Stat.Name, stat.BaseStat})
		}

		for _, past_types := range pokemonApiRes.PastTypes {
			for _, past_type := range past_types.Types {
				pokemonInfo.PokemonTypes = append(pokemonInfo.PokemonTypes, past_type.Type.Name)
			}
		}

		for _, ptype := range pokemonApiRes.Types {
			pokemonInfo.PokemonTypes = append(pokemonInfo.PokemonTypes, ptype.Type.Name)
		}

		return pokemonInfo, nil

	}

	// fmt.Println("Cache Miss")

	req, err := http.NewRequest("GET", fullurl, nil)

	if err != nil {
		return PokemonInfoDefn{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfoDefn{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println(fullurl)
		return PokemonInfoDefn{}, fmt.Errorf("error: returned with staus code %v", res.StatusCode)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfoDefn{}, err
	}

	pokemonApiRes := PokemonInfoAPIDefn{}
	err = json.Unmarshal(body, &pokemonApiRes)
	if err != nil {
		return PokemonInfoDefn{}, err
	}

	numberOfPokemonPastTypes := 0
	for _, past_types := range pokemonApiRes.PastTypes {
		numberOfPokemonPastTypes += len(past_types.Types)
	}

	pokemonInfo := PokemonInfoDefn{
		Name:           pokemonApiRes.Name,
		BaseExperience: pokemonApiRes.BaseExperience,
		Height:         pokemonApiRes.Height,
		Weight:         pokemonApiRes.Weight,
		Stats: make([]struct {
			Name     string
			BaseStat int
		}, 0, len(pokemonApiRes.Stats)),
		PokemonTypes: make([]string, 0, len(pokemonApiRes.Types)+numberOfPokemonPastTypes),
	}

	for _, stat := range pokemonApiRes.Stats {
		pokemonInfo.Stats = append(pokemonInfo.Stats, struct {
			Name     string
			BaseStat int
		}{stat.Stat.Name, stat.BaseStat})
	}

	for _, past_types := range pokemonApiRes.PastTypes {
		for _, past_type := range past_types.Types {
			pokemonInfo.PokemonTypes = append(pokemonInfo.PokemonTypes, past_type.Type.Name)
		}
	}

	for _, ptype := range pokemonApiRes.Types {
		pokemonInfo.PokemonTypes = append(pokemonInfo.PokemonTypes, ptype.Type.Name)
	}

	c.cache.Add(fullurl, body)

	return pokemonInfo, nil

}
