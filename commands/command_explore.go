package commands

import (
	"errors"
	"fmt"
)

func commandExplore(cfig *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	locationResp, err := cfig.PokeapiClient.GetLocationData(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationResp.Name)
	fmt.Println("Found Pokemon:")
	for _, poke := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", poke.Pokemon.Name)
	}

	return nil
}
