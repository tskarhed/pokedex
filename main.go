package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tskarhed/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	PokeAPI *pokeapi.Client
	Pokedex map[string]pokeapi.Pokemon
}

var commands map[string]cliCommand

func main() {

	config := Config{
		PokeAPI: pokeapi.NewClient(),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show all commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show next location areas",
			callback:    commandNextLocationAreas,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous location areas",
			callback:    commandPreviousLocationAreas,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all the pokemon you have caught",
			callback:    commandPokedex,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {

		text := scanner.Text()

		words := cleanInput(text)

		if len(words) == 0 {
			fmt.Printf("Pokedex > ")
			continue
		}

		if command, ok := commands[words[0]]; !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&config, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Printf("Pokedex > ")
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}

func commandExit(config *Config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Printf("\n\n")
	fmt.Println("Available commands:")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandInspect(config *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("inspect requires a valid pokemon name")
	}

	if pokemon, ok := config.Pokedex[args[0]]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf(" - %s\n", t.Type.Name)
		}
		return nil
	} else {
		return fmt.Errorf("you have not caught that pokemon")
	}

}

func commandPokedex(config *Config, _ ...string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("You have not caught any pokemon yet")
		return nil
	}

	fmt.Println("Pokedex:")

	for _, pokemon := range config.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
