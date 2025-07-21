package main

import "fmt"

func commandHelp(config *config, param string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	} // for
	fmt.Println()
	return nil
} // commandHelp
