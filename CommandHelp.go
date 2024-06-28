package main

import "fmt"

func helpCommand(cfg *config) error {
	fmt.Println("You are in help Menu")
	commandList2 := getCommands()
	fmt.Println("Available commands:")
	for _, val := range commandList2 {
		fmt.Printf(" - %v : %v\n", val.name, val.description)
	}
	return nil
}
