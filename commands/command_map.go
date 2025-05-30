package commands

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *Config, args ...string) error {
	err := commandMap(cfg, cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	err := commandMap(cfg, cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	return nil
}

func commandMap(cfg *Config, url *string) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
