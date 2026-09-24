package main

import "fmt"

func merchant(p *Character) {
	for {
		printHeader("Marche noir // Atelier cyber")
		fmt.Println("Credits disponibles :", p.Argent)
		fmt.Println()
		printOption(1, "Stimpack - 3 credits")
		printOption(2, "Cyber Virus - 6 credits")
		printOption(3, "Programme : Surcharge plasma - 25 credits")
		printOption(4, "Puce de donnees - 4 credits")
		printOption(5, "Peau de cyborg - 7 credits")
		printOption(6, "Cuir synthetique - 3 credits")
		printOption(7, "Fibre optique - 1 credit")
		printOption(8, "Batterie - 5 credits")
		printOption(9, "Extension d'inventaire - 30 credits")
		printOption(10, "Retour")
		var choice int
		askChoice()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventaire plein.")
			} else if p.Argent < 3 {
				fmt.Println("Credits insuffisants.")
			} else {
				p.Argent -= 3
				addInventory(p, "Stimpack")
				fmt.Println("Achat confirme : Stimpack.")
			}
		case 2:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventaire plein.")
			} else if p.Argent < 6 {
				fmt.Println("Credits insuffisants.")
			} else {
				p.Argent -= 6
				addInventory(p, "Cyber Virus")
				fmt.Println("Achat confirme : Cyber Virus.")
			}
		case 3:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventaire plein.")
			} else if p.Argent < 25 {
				fmt.Println("Credits insuffisants.")
			} else {
				p.Argent -= 25
				addInventory(p, "Programme : Surcharge plasma")
				fmt.Println("Achat confirme : Programme : Surcharge plasma.")
			}
		case 4:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventaire plein.")
			} else if p.Argent < 4 {
				fmt.Println("Credits insuffisants.")
			} else {
				p.Argent -= 4
				addInventory(p, "Puce de données")
				fmt.Println("Achat confirme : Puce de donnees.")
			}
		case 5:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventaire plein.")
			} else if p.Argent < 7 {
				fmt.Println("Credits insuffisants.")
			} else {
				p.Argent -= 7
				addInventory(p, "Peau de cyborg")
				fmt.Println("Achat confirme : Peau de cyborg.")
			}
		case 6:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventaire plein.")
			} else if p.Argent < 3 {
				fmt.Println("Credits insuffisants.")
			} else {
				p.Argent -= 3
				addInventory(p, "Cuir synthétique")
				fmt.Println("Achat confirme : Cuir synthetique.")
			}
		case 7:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventaire plein.")
			} else if p.Argent < 1 {
				fmt.Println("Credits insuffisants.")
			} else {
				p.Argent -= 1
				addInventory(p, "Fibre optique")
				fmt.Println("Achat confirme : Fibre optique.")
			}
		case 8:
			if len(p.Inventory) >= p.MaxInventory {
				fmt.Println("Inventaire plein.")
			} else if p.Argent < 5 {
				fmt.Println("Credits insuffisants.")
			} else {
				p.Argent -= 5
				addInventory(p, "Batterie")
				fmt.Println("Achat confirme : Batterie.")
			}
		case 9:
			upgradeInventorySlot(p)
		case 10:
			return
		default:
			fmt.Println("\nChoix invalide.")
		}
	}
}
