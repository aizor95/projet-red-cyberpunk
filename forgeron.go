package main

import "fmt"

func forgeron(p *Character) {
	fmt.Println("\n===================================")
	fmt.Println("             FORGERON              ")
	fmt.Println("===================================")
	fmt.Println("1. Visière Tactique (5 pièces d'or)")
	fmt.Println("   - Requis : 1 Puce de Données, 1 Cuir Synthétique")
	fmt.Println("2. Veste en Fibre de Carbone (5 pièces d'or)")
	fmt.Println("   - Requis : 2 Fibres Optiques, 1 Peau de Cyborg")
	fmt.Println("3. Bottes Cybernétiques (5 pièces d'or)")
	fmt.Println("   - Requis : 1 Fibre Optique, 1 Cuir Synthétique")
	fmt.Println("4. Retour au menu principal")
	
	var choice int
	fmt.Print("\nVotre choix : ")
	fmt.Scanln(&choice)
	
	cost := 5

	countItem := func(itemName string) int {
		count := 0
		for _, item := range p.Inventory {
			if item == itemName {
				count++
			}
		}
		return count
	}

	removeItem := func(itemName string) {
		for i, item := range p.Inventory {
			if item == itemName {
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
				break
			}
		}
	}

	switch choice {
	case 1:
		itemToCraft := "Visière Tactique"
		mat1 := "Puce de Données"
		mat2 := "Cuir Synthétique"

		if p.Argent < cost {
			fmt.Println("\nErreur : Fonds insuffisants ! Vous avez besoin de 5 pièces d'or.")
			return
		}
		if countItem(mat1) < 1 || countItem(mat2) < 1 {
			fmt.Println("\nErreur : Matériaux manquants pour fabriquer la Visière Tactique (1 Puce de Données et 1 Cuir Synthétique requis).")
			return
		}

		p.Argent -= cost
		removeItem(mat1)
		removeItem(mat2)
		
		p.Inventory = append(p.Inventory, itemToCraft)
		fmt.Println("\nSuccès : Visière Tactique fabriquée et ajoutée à votre inventaire !")

	case 2:
		itemToCraft := "Veste en Fibre de Carbone"
		mat1 := "Fibre Optique"
		mat2 := "Peau de Cyborg"

		if p.Argent < cost {
			fmt.Println("\nErreur : Fonds insuffisants ! Vous avez besoin de 5 pièces d'or.")
			return
		}
		if countItem(mat1) < 2 || countItem(mat2) < 1 {
			fmt.Println("\nErreur : Matériaux manquants pour fabriquer la Veste en Fibre de Carbone (2 Fibres Optiques et 1 Peau de Cyborg requis).")
			return
		}

		p.Argent -= cost
		removeItem(mat1)
		removeItem(mat1)
		removeItem(mat2)
		
		p.Inventory = append(p.Inventory, itemToCraft)
		fmt.Println("\nSuccès : Veste en Fibre de Carbone fabriquée et ajoutée à votre inventaire !")

	case 3:
		itemToCraft := "Bottes Cybernétiques"
		mat1 := "Fibre Optique"
		mat2 := "Cuir Synthétique"

		if p.Argent < cost {
			fmt.Println("\nErreur : Fonds insuffisants ! Vous avez besoin de 5 pièces d'or.")
			return
		}
		if countItem(mat1) < 1 || countItem(mat2) < 1 {
			fmt.Println("\nErreur : Matériaux manquants pour fabriquer les Bottes Cybernétiques (1 Fibre Optique et 1 Cuir Synthétique requis).")
			return
		}

		p.Argent -= cost
		removeItem(mat1)
		removeItem(mat2)
		
		p.Inventory = append(p.Inventory, itemToCraft)
		fmt.Println("\nSuccès : Bottes Cybernétiques fabriquées et ajoutées à votre inventaire !")

	case 4:
		fmt.Println("\nRetour au menu principal...")
	default:
		fmt.Println("\nErreur : Choix invalide.")
	}
}