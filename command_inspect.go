package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]
	if pokemon, ok := cfg.caughtPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s\n", pokemonName)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, v := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Types:")
		for _, v := range pokemon.Types {
			fmt.Printf("  - %s\n", v.Type.Name)
		}
		return nil
	}
	return fmt.Errorf("you have not caught that pokemon")
}
