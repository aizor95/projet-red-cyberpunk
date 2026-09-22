package main

import "fmt"

// MARCHAND

func marchand(p *Character) {
	for {
		fmt.Println("\n===================================")
		fmt.Println("       CYBERNETIC WORKSHOP")
		fmt.Println("===================================")
		fmt.Println("Your credits:", p.Argent)
		fmt.Println()
		fmt.Println("1. Stimpack (3 credits)")
		fmt.Println("2. Cyber Virus (6 credits)")
		fmt.Println("3. Spellbook: Boule de Feu (25 credits)")
		fmt.Println("4. Fourrure de Loup (4 credits)")
		fmt.Println("5. Peau de Troll (7 credits)")
		fmt.Println("6. Cuir de Sanglier (3 credits)")
		fmt.Println("7. Plume de Corbeau (1 credits)")
		fmt.Println("8. Augmentation d' inventaire(30 credits)")
		fmt.Println("9. Back")

		var choice int
		fmt.Print("\nYour choice: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 3 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 3
				addInventory(p, "Stimpack")
				fmt.Println("You bought Stimpack")
			}

		case 2:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 6 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 6
				addInventory(p, "Cyber Virus")
				fmt.Println("You bought Cyber Virus")
			}

		case 3:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 25 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 25
				addInventory(p, "Spellbook: Boule de Feu")
				fmt.Println("You bought Spellbook: Boule de Feu")
			}

		case 4:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 4 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 4
				addInventory(p, "Fourrure de Loup")
				fmt.Println("You bought Fourrure de Loup")
			}

		case 5:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 7 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 7
				addInventory(p, "Peau de Troll")
				fmt.Println("You bought Peau de Troll")
			}

		case 6:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 3 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 3
				addInventory(p, "Cuir de Sanglier")
				fmt.Println("You bought Cuir de Sanglier")
			}

		case 7:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 1 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 1
				addInventory(p, "Plume de Corbeau")
				fmt.Println("You bought Plume de Corbeau")
			}

		case 8:
			upgradeInventorySlot(p)

		case 9:
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
