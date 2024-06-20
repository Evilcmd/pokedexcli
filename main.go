package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type commandDetails struct {
	name        string
	description string
	callback    func() error
}

type locationMapResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var (
	errExit      = errors.New("exit pokedex")
	baseMapUrl   = "https://pokeapi.co/api/v2/location-area/?offset="
	mapUrlOffset = -20
	mp           = map[string]commandDetails{
		"help": {"help", "Displays help message", nil},
		"exit": {"exit", "Exit the Pokedex", commandExit},
		"map":  {"map", "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations", commandMap},
		"mapb": {"mapb", "displays the previous 20 locations, If you're on the first page of results, it displays an error message.", commandMapb},
	}
)

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, v := range mp {
		fmt.Println(v.name + ": " + v.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	return errExit
}

func getLocations() (locationMapResponse, error) {
	res, err := http.Get(baseMapUrl + strconv.Itoa(mapUrlOffset) + "&limit=20")
	if err != nil {
		return locationMapResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return locationMapResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return locationMapResponse{}, err
	}
	locationMap := locationMapResponse{}
	err = json.Unmarshal(body, &locationMap)
	if err != nil {
		return locationMapResponse{}, err
	}
	if len(locationMap.Results) == 0 {
		return locationMapResponse{}, errors.New("no more loctions available")
	}
	return locationMap, nil
}

func commandMap() error {
	mapUrlOffset += 20
	locationMap, err := getLocations()
	if err != nil {
		return nil
	}
	for _, v := range locationMap.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandMapb() error {
	if mapUrlOffset-20 < 0 {
		return errors.New("cannot go back a page")
	}
	mapUrlOffset -= 20
	locationMap, err := getLocations()
	if err != nil {
		return nil
	}
	for _, v := range locationMap.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func main() {

	for {
		fmt.Print("pokedex > ")
		var command string
		fmt.Scan(&command)
		if command == "help" {
			commandHelp()
		} else if commStruct, ok := mp[command]; ok {
			if x := commStruct.callback(); x != nil {
				if x == errExit {
					return
				}
				fmt.Println("Error: " + x.Error())
			}
		} else {
			fmt.Println("Command not found")
		}
	}
}
