package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		lowerCaseUserInput := strings.ToLower(userInput)
		words := strings.Fields(lowerCaseUserInput)
		fmt.Printf("Your command was: %s\n", words[0])
	}
}
