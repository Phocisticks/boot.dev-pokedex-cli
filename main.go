package main

import (
	"time"

	"github.com/phocisticks/pokedexcli/commands"
	"github.com/phocisticks/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &commands.Config{
		PokeapiClient: pokeClient,
		Pokedex:       map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
