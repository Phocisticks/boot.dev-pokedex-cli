package commands

import (
	"errors"
	"fmt"
	"math/rand"
)

const max_chance = 400

func commandCatch(cfig *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemonResp, err := cfig.PokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	chance := pokemonResp.BaseExperience

	attempt := rand.Intn(max_chance)

	if attempt >= chance {
		cfig.Pokedex[name] = pokemonResp

		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
