package commands

import (
	"fmt"
	"github.com/interyx/pokedexcli/pokeapi"
	"os"
)

type Config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config, input string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Display information about the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display information about the previous 20 locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore <area>",
			Description: "Explore the given area to find all the Pokemon living there\nFind area names with 'map'",
			Callback:    commandExplore,
		},
	}
}

func commandHelp(cfg *Config, input string) error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Printf("\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	fmt.Printf("\n")
	return nil
}

func commandExit(cfg *Config, input string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config, input string) error {
	locations, newPrevious, newNext, err := pokeapi.GetNextLocation(&cfg.Next)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Printf("%s\n", location.Name)
	}
	cfg.Next = newNext
	cfg.Previous = newPrevious
	return nil
}

func commandMapb(cfg *Config, input string) error {
	locations, newPrevious, newNext, err := pokeapi.GetNextLocation(&cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = newNext
	cfg.Previous = newPrevious
	for _, location := range locations {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandExplore(cfg *Config, input string) error {
	pokemans, err := pokeapi.GetPokemonAtLocation(input)
	if err != nil {
		return err
	}
	return nil
}
