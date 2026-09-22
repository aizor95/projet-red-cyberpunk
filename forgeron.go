package main

import "fmt"

func forgeron(p *Character) {
	fmt.Println("\n===================================")
	fmt.Println("            FORGERON             ")
	fmt.Println("===================================")
	fmt.Println("1. Chapeau de l'aventurier (Coût : 5 Credit)")
	fmt.Println("2. Tunique de l'aventurier (Coût : 5 Credit)")
	fmt.Println("3. Bottes de l'aventurier (Coût : 5 Credit)")
	fmt.Println("4. Retour")

	var choice int
	fmt.Print("\nVotre choix : ")
	fmt.Scanln(&choice)

	cost := 5

	switch choice {
	case 1, 2, 3:
		if p.Argent >= cost {
			p.Argent -= cost
			var itemName string

			switch choice {
			case 1:
				itemName = "Chapeau de l'aventurier"
			case 2:
				itemName = "Tunique de l'aventurier"
			case 3:
				itemName = "Bottes de l'aventurier"
			}

			p.Inventory = append(p.Inventory, itemName)

			fmt.Println("\nObjet fabriqué avec succès !")
			fmt.Println("Vous avez obtenu :", itemName)
			fmt.Println("Solde restant :", p.Argent, "Credit.")
		} else {
			fmt.Println("\nErreur : Vous n'avez pas assez d'or ! (Il vous faut 5 pièces).")
		}
	case 4:
		fmt.Println("\nRetour au menu principal...")
	default:
		fmt.Println("\nChoix invalide.")
	}
}
