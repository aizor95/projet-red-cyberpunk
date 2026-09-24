package main

import (
	"fmt"
	"strings"
)

func printHeader(title string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 42))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 42))
}

func printPanel(title string) {
	fmt.Println()
	fmt.Println(strings.Repeat("-", 42))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("-", 42))
}

func printOption(number int, label string) {
	fmt.Printf("  %d. %s\n", number, label)
}

func askChoice() {
	fmt.Print("\n> Votre choix : ")
}
