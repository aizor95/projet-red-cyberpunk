package main

import "fmt"

// MENU PRINCIPAL

func mainMenu(p *Character) {
	for {
		fmt.Println("\n===================================")
		fmt.Println("            NEON CITY")
		fmt.Println("===================================")
		fmt.Println("1. Display character profile")
		fmt.Println("2. Access inventory")
		fmt.Println("3. Cybernetic Workshop")
		fmt.Println("4. Training")
		fmt.Println("5. Quit")
		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			displayInfo(p)
		case 2:
			inventoryMenu(p)
		case 3:
			merchant(p)
		case 4:
			trainingFight(p)
		case 5:
			fmt.Println("\nExiting Neon City system...")
			return
		default:
			fmt.Println("\nInvalid choice, please try again.")
		}
	}
}
