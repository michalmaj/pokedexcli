package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(strings.TrimSpace(text))
	if text == "" {
		return []string{}
	}
	return strings.Fields(text)
}

func main() {
	fmt.Println("Hello, World!")
}