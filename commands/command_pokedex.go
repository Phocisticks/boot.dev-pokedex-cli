package commands

import "fmt"

func commandPokedex(cfig *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, poke := range cfig.Pokedex {
		fmt.Printf("\t- %s\n", poke.Name)
	}

	return nil
}
