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
		inputTokens := strings.Split(input, " ")
		cmd := strings.ToLower(strings.Trim(inputTokens[0], "\n"))
		subject := ""
		if len(inputTokens) > 1 {
			subject = strings.ToLower(strings.Trim(inputTokens[1], "\n"))
		}
		commands := commands.GetCommands()
		command, ok := commands[cmd]
		if !ok {
			fmt.Println("That command is not recognized.  If you need help, try 'help'.")
		}
		if err = command.Callback(&cfg, subject); err != nil {
			fmt.Println("An error has occurred: %w", err)
		}
	}
}
