package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")

		scanSuccessful := scanner.Scan()

		if !scanSuccessful {
			break
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("reading input:", err)
			break
		}
		words := cleanInput(scanner.Text())
		if len(words) <= 0 {
			fmt.Println("Invalid input, try again")
			continue
		}

		commandName := words[0]

		if cmd, exists := getCommands()[commandName]; exists {
			if err := cmd.callback(); err != nil {
				fmt.Println("err")
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    	commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

