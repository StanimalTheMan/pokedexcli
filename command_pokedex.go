package main

import "fmt"

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemonName, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
