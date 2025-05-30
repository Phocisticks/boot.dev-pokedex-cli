package commands

import (
	"errors"
	"fmt"

	"github.com/phocisticks/pokedexcli/internal/pokeapi"
)

func commandInspect(cfig *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, exists := cfig.Pokedex[name]
	if !exists {
		fmt.Printf("you have not caught %s \n", name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	printPokemonStats(pokemon)

	return nil
}

func printPokemonStats(poke pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %v\n", poke.Height)
	fmt.Printf("Weight: %v\n", poke.Weight)
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf("\t-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}
}
