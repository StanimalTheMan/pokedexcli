package main

import (
	"time"

	"github.com/StanimalTheMan/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeapi.NewClient(5*time.Second, time.Minute*5),
	}

	startRepl(&cfg)
}
