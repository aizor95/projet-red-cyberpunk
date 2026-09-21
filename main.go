package main

import "fmt"

func main() {
	fmt.Println("===================================")
	fmt.Println("            NEON CITY")
	fmt.Println("===================================")
	fmt.Println("\nWelcome to Neon City, year 2087.")
	fmt.Println("Create your cybernetic character.\n")

	character := characterCreation()

	fmt.Println("\n===================================")
	fmt.Println("       CHARACTER CREATED")
	fmt.Println("===================================")
	fmt.Println("Welcome,", character.Name)
	fmt.Println("Class:", character.Class)

	mainMenu(&character)
}
