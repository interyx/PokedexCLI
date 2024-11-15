package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("pokedex > ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input")
			return
		}
		input = strings.TrimSuffix(input, "\n")
		fmt.Println(input)
	}
}
