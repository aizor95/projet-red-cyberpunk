package main

import "fmt"

// MENU PRINCIPAL

func mainMenu(p *Character) {
	for {
		printHeader("Tableau de bord")
		fmt.Println("Position :", p.CurrentLocation)
		fmt.Println("Credits  :", p.Argent)
		fmt.Printf("Etat     : %d/%d HP | %d/%d energie\n", p.HP, p.MaxHP, p.Energy, p.MaxEnergy)
		fmt.Println()
		printOption(1, "Profil du personnage")
		printOption(2, "Ouvrir l'inventaire")
		printOption(3, "Explorer Neon City")
		printOption(4, "Consulter les missions")
		printOption(5, "Gerer l'equipement")
		printOption(6, "Quitter")
		var choice int
		askChoice()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			displayInfo(p)
		case 2:
			inventoryMenu(p)
		case 3:
			explore(p)
		case 4:
			missions(p)
		case 5:
			EquipeArmor(p)
		case 6:
			fmt.Println("\nDeconnexion du reseau de Neon City...")
			return
		default:
			fmt.Println("\nCommande inconnue.")
		}
	}
}
