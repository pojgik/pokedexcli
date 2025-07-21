package main

import "fmt"

func commandPokedex(config *config, param string) error {
	fmt.Println("Your Pokedex:")
	for k := range config.caught {
		fmt.Println(" - ", k)
	} // for
	return nil
} // commandPokedex
