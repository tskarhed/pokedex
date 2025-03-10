package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("catch requires a valid pokemon name")
	}
	pokemonName := args[0]

	pokemon, err := config.PokeAPI.GetPokemon(pokemonName)
	if err != nil {
		fmt.Println("Error fetching pokemon:")
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")

	baseExperience := pokemon.BaseExperience
	random := rand.Intn(baseExperience)

	if random > baseExperience/2 {
		fmt.Println("You caught " + pokemonName + "!")
	} else {
		fmt.Println("You missed " + pokemonName + "!")
		return nil
	}

	config.Pokedex[pokemonName] = pokemon

	return nil
}
