package main

import "fmt"

var countKills int
var missionKillActive bool

func missions(p *Character) {
	fmt.Println("\n===================================")
	fmt.Println("        AVAILABLE MISSIONS")
	fmt.Println("===================================")
	fmt.Println("1. Kill 3 monsters (+3 CREDITS)")
	fmt.Println("2. Data Delivery (+3 CREDITS)")
	fmt.Println("3. Terminal Hack (+3 CREDITS)")
	var mission int
	fmt.Print("\nEnter your choice: ")
	fmt.Scanln(&mission)
	switch mission {
	case 1:
		if !missionKillActive {
			fmt.Println("\nMission : Kill 3 monsters")
			fmt.Println("Reward : +3 CREDITS")
			fmt.Println("1. Accept mission")
			fmt.Println("2. Back")
			var accept int
			fmt.Print("\nEnter your choice: ")
			fmt.Scanln(&accept)
			if accept == 1 {
				missionKillActive = true
				countKills = 0
				fmt.Println("\nMission accepted!")
				fmt.Println("You need to kill 3 monsters.")
			} else {
				fmt.Println("\nMission not accepted.")
			}
		} else {
			fmt.Printf("\nMonsters killed: %d/3\n", countKills)
			if countKills >= 3 {
				p.Argent += 3
				fmt.Println("\nMission completed!")
				fmt.Println("+3 crédits.")
				fmt.Println("TOTAL :", p.Argent, "CREDIT")
				countKills = 0
				missionKillActive = false
			} else {
				fmt.Printf("You need to kill %d more monster(s).\n", 3-countKills)
			}
		}
	case 2:
		p.Argent += 3
		fmt.Println("\nLivraison effectuée avec succès !")
		fmt.Println("+3 crédits.")
		fmt.Println("TOTAL :", p.Argent, "CREDIT")
	case 3:
		p.Argent += 3
		fmt.Println("\nTerminal piraté avec succès !")
		fmt.Println("+3 crédits.")
		fmt.Println("TOTAL :", p.Argent, "CREDIT")
	default:
		fmt.Println("\nChoix invalide.")
	}
}