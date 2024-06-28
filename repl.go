package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			"help",
			"help menu",
			helpCommand,
		},
		"exit": {
			"exit",
			"exit the program",
			exitCommand,
		},
		"map": {
			"map",
			"Lists some location areas",
			mapCommand,
		},
		"mapb": {
			"mapb",
			"Lists location areas of the previous page",
			mapbCommand,
		},
		"explore": {
			"explore",
			"Lists all the pokemons in the area",
			exploreCommand,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		commandList := getCommands()
		command, ok := commandList[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		args := ""
		if len(cleaned) > 1 {
			args = cleaned[1]
			fmt.Println(args)
		}
		err := command.callback(cfg, args)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(inputString string) (splitWords []string) {
	lowered := strings.ToLower(inputString)
	splitWords = strings.Fields(lowered)
	return
}
