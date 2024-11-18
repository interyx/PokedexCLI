package commands

import (
	"fmt"
	"os"

	"github.com/interyx/pokedexcli/pokeapi"
)

type Config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config) error
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
	}
}

func commandHelp(cfg *Config) error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Printf("\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	fmt.Printf("\n")
	return nil
}

func commandExit(cfg *Config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	locations, newPrevious, newNext := pokeapi.GetNextLocation(&cfg.Next)
	for _, location := range locations {
		fmt.Printf("%s\n", location.Name)
	}
	cfg.Next = newNext
	cfg.Previous = newPrevious
	return nil
}

func commandMapb(cfg *Config) error {
	locations, newPrevious, newNext := pokeapi.GetNextLocation(&cfg.Previous)
	cfg.Next = newNext
	cfg.Previous = newPrevious
	for _, location := range locations {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
