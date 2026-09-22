package main

import "fmt"

func missions(p *Character) {
	var countkills int
	fmt.Println("\n===================================")
	fmt.Printf("AVAILABLE MISSIONS")
	fmt.Println("\n===================================")
	fmt.Println("1. Kill 3 monsters (+3 CREDITS)")
	fmt.Println("2. Data Delivery (+3 CREDITS)")
	fmt.Println("3. Terminal Hack (+3 CREDITS)")
	var mission int
	fmt.Print("\nEnter your choice: ")
	fmt.Scanln(&mission)
	switch mission {
	case 1:
		if countkills == 3 {
			p.Argent += 3
			fmt.Println("+3 crédits. Solde total :", p.Argent, "Euro.")
			countkills = 0
		} else {
			println("Monsters killed: ", countkills)
			fmt.Printf("You need to kill %d monsters", (3 - countkills))
		}
	case 2:
		p.Argent += 3
		fmt.Println("\nLivraison effectuée avec succès !")
		fmt.Println("+3 crédits.")
		fmt.Println("TOTAL :", p.Argent, " CREDIT ")
	case 3:
		p.Argent += 3
		fmt.Println("\nTerminal piraté avec succès !")
		fmt.Println("+3 crédits.")
		fmt.Println("TOTAL :", p.Argent, " CREDIT ")
	default:
		fmt.Println("\nChoix invalide.")
	}
}
