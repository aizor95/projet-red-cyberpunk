package main

import "fmt"

var countKills int
var missionKillActive bool

func readIntInput() int {
	var input string
	fmt.Scanln(&input)
	isNumber := true
	for i := 0; i < len(input); i++ {
		if input[i] < '0' || input[i] > '9' {
			isNumber = false
			break
		}
	}
	if !isNumber || len(input) == 0 {
		for i := 0; i < len(input); i++ {
			fmt.Println("Erreur : Veuillez entrer un nombre valide !")
		}
		return -1
	}
	val := 0
	for i := 0; i < len(input); i++ {
		val = val*10 + int(input[i]-'0')
	}
	return val
}

func missions(p *Character) {
	for {
		fmt.Println("\n===================================")
		fmt.Println("        AVAILABLE MISSIONS")
		fmt.Println("===================================")
		if missionKillActive {
			fmt.Printf("1. Kill 3 monsters (+3 CREDITS) [EN COURS: %d/3]\n", countKills)
		} else {
			fmt.Println("1. Kill 3 monsters (+3 CREDITS)")
		}
		fmt.Println("2. Data Delivery (+3 CREDITS)")
		fmt.Println("3. Terminal Hack (+3 CREDITS)")
		fmt.Println("0. Back")
		fmt.Print("\nEnter your choice: ")
		mission := readIntInput()
		if mission == 0 {
			break
		}
		switch mission {
		case 1:
			if !missionKillActive {
				fmt.Println("\nMission : Kill 3 monsters")
				fmt.Println("Reward : +3 CREDITS")
				fmt.Println("1. Accept mission")
				fmt.Println("2. Back")
				fmt.Print("\nEnter your choice: ")
				accept := readIntInput()
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
			if missionKillActive {
				fmt.Println("\n/!\\ ATTENTION /!\\")
				fmt.Println("Vous avez deja une quete en cours !")
				fmt.Println("Si vous changez de quete, vous allez perdre votre progression.")
				fmt.Print("Voulez-vous continuer ? (1 = Oui / 2 = Non) : ")
				choix := readIntInput()
				if choix != 1 {
					fmt.Println("\nChangement de quete annule.")
					continue
				}
				missionKillActive = false
				countKills = 0
			}
			p.Argent += 3
			fmt.Println("\nLivraison effectuée avec succès !")
			fmt.Println("+3 crédits.")
			fmt.Println("TOTAL :", p.Argent, "CREDIT")
		case 3:
			if missionKillActive {
				fmt.Println("\n/!\\ ATTENTION /!\\")
				fmt.Println("Vous avez deja une quete en cours !")
				fmt.Println("Si vous changez de quete, vous allez perdre votre progression.")
				fmt.Print("Voulez-vous continuer ? (1 = Oui / 2 = Non) : ")
				choix := readIntInput()
				if choix != 1 {
					fmt.Println("\nChangement de quete annule.")
					continue
				}
				missionKillActive = false
				countKills = 0
			}
			p.Argent += 3
			fmt.Println("\nTerminal piraté avec succès !")
			fmt.Println("+3 crédits.")
			fmt.Println("TOTAL :", p.Argent, "CREDIT")
		default:
			if mission != -1 {
				fmt.Println("\nChoix invalide.")
			}
		}
	}
}