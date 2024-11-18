package main

import (
	"bufio"
	"fmt"
	"github.com/interyx/pokedexcli/commands"
	"os"
	"strings"
)

func main() {
	cfg := commands.Config{}
	for {
		fmt.Print("pokedex > ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input: %w", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		input = strings.ToLower(input)
		commands := commands.GetCommands()
		command, ok := commands[input]
		if !ok {
			fmt.Println("That command is not recognized.  If you need help, try 'help'.")
		}
		if err = command.Callback(&cfg); err != nil {
			fmt.Println("An error has occurred: %w", err)
		}
	}
}
