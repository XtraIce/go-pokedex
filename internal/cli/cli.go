package cli

import (
	"fmt"
	"sort"
)

// CliCommand represents a command in the CLI.
type CliCommand struct {
	Name        string
	Description string
	Callback    func(args []string) error
}

// CliCommands is a map of command names to CliCommand objects.
var CliCommands = map[string]CliCommand{}

func init() {
	CliCommands = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback: func([]string) error {
				fmt.Println("Goodbye!")
				return fmt.Errorf("exit")
			},
		},
		"help": {
			Name:        "help",
			Description: "Show help",
			Callback: func([]string) error {
				keys := make([]string, 0, len(CliCommands))
				for name, _ := range CliCommands {
					keys = append(keys, name)
				}
				sort.Strings(keys)
				for _, name := range keys {
					command := CliCommands[name]
					fmt.Println(command.Name, ":", command.Description)
				}
				return nil
			},
		},
		"map": {
			Name:        "map",
			Description: "Displays Names of 20 location areas in Pokemon world",
			Callback: func([]string) error {
				getLocations()
				return nil
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays Names of 20 previous location areas in Pokemon world",
			Callback: func([]string) error {
				getLocationsBefore()
				return nil
			},
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location area",
			Callback: func(args []string) error {
				if len(args) < 2 {
					fmt.Println("Please provide a location name")
					return nil
				}
				exploreLocation(args[1])
				return nil
			},
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback: func(args []string) error {
				if len(args) < 2 {
					fmt.Println("Please provide a pokemon name")
					return nil
				}
				PokemonCatch(args[1])
				return nil
			},
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a caught pokemon",
			Callback: func(args []string) error {
				if len(args) < 2 {
					fmt.Println("Please provide a pokemon name")
					return nil
				}
				PokemonInspect(args[1])
				return nil
			},
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List caught pokemon",
			Callback: func([]string) error {
				PokedexList()
				return nil
			},
		},
	}
}
