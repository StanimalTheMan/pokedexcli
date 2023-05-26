package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name	    string
	description string
	callback    func() error
}

func commandHelp () error {
	fmt.Println("Welcome to the Pokedex!")
}

func commandExit() error {
	return nil
}

func getCliCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:	     "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		userInput := scanner.Text()
		if userInput == "help" {
			getCliCommand()["help"]
		} else if userInput == "exit" {
			getCliCommand()["exit"]
		}
	}
}