package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
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

func commandHelp() error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Printf("\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	fmt.Printf("\n")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	fmt.Println("This is where the map goes.")
	return nil
}

func commandMapb() error {
	fmt.Println("Previous maps go here")
	return nil
}
