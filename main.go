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
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
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
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		if userInput == "help" {
			fmt.Print(getCliCommand()["help"].callback())
		} else if userInput == "exit" {
			fmt.Print(getCliCommand()["exit"].callback())
		}
	}
}