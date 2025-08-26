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
		fmt.Print("Pokedex >")
		scanSuccessful := scanner.Scan()
		if !scanSuccessful {
			fmt.Printf("exit\nQuitting program")
			break
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("reading input:", err)
			break
		}
		userText := scanner.Text()
		words := cleanInput(userText)
		if len(words) > 0 {
			fmt.Println("Your command was:", words[0])
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
