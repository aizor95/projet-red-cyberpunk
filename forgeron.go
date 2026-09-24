package main

import "fmt"

func forgeron(p *Character) {
	printHeader("Marche noir // Forgeron")
	fmt.Println("Credits disponibles :", p.Argent)
	fmt.Println()
	printOption(1, "Visiere tactique - 5 credits")
	fmt.Println("     Requis : 1 Puce de donnees, 1 Cuir synthetique")
	printOption(2, "Veste en fibre de carbone - 5 credits")
	fmt.Println("     Requis : 2 Fibres optiques, 1 Peau de cyborg")
	printOption(3, "Bottes cybernetiques - 5 credits")
	fmt.Println("     Requis : 1 Fibre optique, 1 Cuir synthetique")
	printOption(4, "Retour")
	var choice int
	askChoice()
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
		itemToCraft := "Visière tactique"
		mat1 := "Puce de données"
		mat2 := "Cuir synthétique"
		if p.Argent < cost {
			fmt.Println("\nCredits insuffisants. Il faut 5 credits.")
			return
		}
		if countItem(mat1) < 1 || countItem(mat2) < 1 {
			fmt.Println("\nMateriaux manquants : 1 Puce de donnees et 1 Cuir synthetique requis.")
			return
		}
		p.Argent -= cost
		removeItem(mat1)
		removeItem(mat2)
		p.Inventory = append(p.Inventory, itemToCraft)
		fmt.Println("\nFabrication terminee : Visiere tactique ajoutee a l'inventaire.")
	case 2:
		itemToCraft := "Veste en fibre de carbone"
		mat1 := "Fibre optique"
		mat2 := "Peau de cyborg"
		if p.Argent < cost {
			fmt.Println("\nCredits insuffisants. Il faut 5 credits.")
			return
		}
		if countItem(mat1) < 2 || countItem(mat2) < 1 {
			fmt.Println("\nMateriaux manquants : 2 Fibres optiques et 1 Peau de cyborg requis.")
			return
		}
		p.Argent -= cost
		removeItem(mat1)
		removeItem(mat1)
		removeItem(mat2)
		p.Inventory = append(p.Inventory, itemToCraft)
		fmt.Println("\nFabrication terminee : Veste en fibre de carbone ajoutee a l'inventaire.")
	case 3:
		itemToCraft := "Bottes cybernétiques"
		mat1 := "Fibre optique"
		mat2 := "Cuir synthétique"
		if p.Argent < cost {
			fmt.Println("\nCredits insuffisants. Il faut 5 credits.")
			return
		}
		if countItem(mat1) < 1 || countItem(mat2) < 1 {
			fmt.Println("\nMateriaux manquants : 1 Fibre optique et 1 Cuir synthetique requis.")
			return
		}
		p.Argent -= cost
		removeItem(mat1)
		removeItem(mat2)

		p.Inventory = append(p.Inventory, itemToCraft)
		fmt.Println("\nFabrication terminee : Bottes cybernetiques ajoutees a l'inventaire.")
	case 4:
		fmt.Println("\nRetour.")
	default:
		fmt.Println("\nChoix invalide.")
	}
}
