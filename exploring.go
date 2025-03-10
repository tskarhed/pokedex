package main

import "fmt"

func commandNextLocationAreas(config *Config, _ ...string) error {

	locationAreas, err := config.PokeAPI.GetNextLocationAreas()
	if err != nil {
		return err
	}

	printLocationAreas(locationAreas.Results)
	return nil
}

func commandPreviousLocationAreas(config *Config, _ ...string) error {
	locationAreas, err := config.PokeAPI.GetPreviousLocationAreas()
	if err != nil {
		return err
	}

	printLocationAreas(locationAreas.Results)
	return nil
}

func commandExplore(config *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("explore requires a valid location area name")
	}
	locationAreaName := args[0]
	locationAreaDetail, err := config.PokeAPI.GetLocationAreaDetail(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + locationAreaName + "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationAreaDetail.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}

	return nil
}

func printLocationAreas(locationAreas []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}) {
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
}
