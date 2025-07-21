package main

import (
	"fmt"
)

func commandInspect(config *config, param string) error {
	pokemon, exists := config.caught[param]
	if !exists {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		} // for

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		} // for
	} // if
	return nil
} // commandInspect
