package commands

import (
	"fmt"
	"os"
)

func commandExit(cfig *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
