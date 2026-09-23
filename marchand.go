package main

import "fmt"

// MARCHAND

func merchant(p *Character) {
	for {
		fmt.Println("\n===================================")
		fmt.Println("       CYBERNETIC WORKSHOP")
		fmt.Println("===================================")
		fmt.Println("Your credits:", p.Argent)
		fmt.Println()
		fmt.Println("1. Stimpack (3 credits)")
		fmt.Println("2. Cyber Virus (6 credits)")
		fmt.Println("3. Surcharge plasma (25 credits)")
		fmt.Println("4. Puce de données (4 crédits)")
		fmt.Println("5. Peau de cyborg (7 crédits)")
		fmt.Println("6. Cuir synthétique (3 crédits)")
		fmt.Println("7. Fibre optique (1 crédit)")
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
				addInventory(p, "Programme : Surcharge plasma")
				fmt.Println("You bought Programme : Surcharge plasma")
			}
		case 4:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 4 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 4
				addInventory(p, "Puce de données")
				fmt.Println("You bought Puce de données")
			}
		case 5:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 7 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 7
				addInventory(p, "Peau de cyborg")
				fmt.Println("You bought Peau de cyborg")
			}
		case 6:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 3 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 3
				addInventory(p, "Cuir synthétique")
				fmt.Println("You bought Cuir synthétique")
			}
		case 7:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventory is full")
			} else if p.Argent < 1 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 1
				addInventory(p, "Fibre optique")
				fmt.Println("You bought Fibre optique")
			}
		case 8:
			upgradeInventorySlot(p)
		case 9:
			return
		default:
			fmt.Println("\nInvalid choice")
		}
	}
}