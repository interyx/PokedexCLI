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
		inputTokens := strings.Split(input, " ")
		if err != nil {
			fmt.Println("An error occurred while reading input: %w", err)
			continue
		}
		if len(inputTokens) > 2 {
			fmt.Println("Warning: too many words detected.")
			fmt.Println("Pokedex currently supports commands in the <command> <subject> format.")
		}
		inputStr := strings.ToLower(inputTokens[0])
		commands := commands.GetCommands()
		command, ok := commands[inputStr]
		if !ok {
			fmt.Println("That command is not recognized.  If you need help, try 'help'.")
		}
		if err = command.Callback(&cfg, strings.ToLower(inputTokens[1])); err != nil {
			fmt.Println("An error has occurred: %w", err)
		}
	}
}
