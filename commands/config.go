package commands

import "github.com/phocisticks/pokedexcli/internal/pokeapi"

type Config struct {
	PokeapiClient    pokeapi.Client
	Pokedex          map[string]pokeapi.Pokemon
	NextLocationsURL *string
	PrevLocationsURL *string
}
