package main

// Import packages
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function that clears input text
// Returns slice with words splitted by space and converted to lower
// Input text is also trimmed from whitespaces
func cleanInput(text string) []string {
	text = strings.ToLower(strings.TrimSpace(text))
	if text == "" {
		return []string{}
	}
	return strings.Fields(text)
}

// Structure that keeps commands blueprint
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Map that contains commands
var commands = map[string]cliCommand{}

// Commands as functions
func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func main() {
	// Insert commands and key-value pairs into map
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println()
			break
		}

		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		cmd, ok := commands[words[0]]
		
		if !ok {
			fmt.Printf("Unknown command")
			continue
		}

		err := cmd.callback()

		if err != nil {
			fmt.Println(err)
		}
	}
}