package main

import "fmt"

// ===================================
// LECTURE D'UN NOMBRE
// ===================================

func readIntInput() int {

	var input string

	fmt.Scanln(&input)

	if len(input) == 0 {
		return -1
	}

	for i := 0; i < len(input); i++ {

		if input[i] < '0' || input[i] > '9' {
			return -1
		}
	}

	value := 0

	for i := 0; i < len(input); i++ {

		value = value*10 + int(input[i]-'0')
	}

	return value
}

// ===================================
// MENU DES QUÊTES
// ===================================

func missions(p *Character) {

	for {

		fmt.Println("\n===================================")
		fmt.Println("          AVAILABLE QUESTS")
		fmt.Println("===================================")

		fmt.Println(
			"Badge level:",
			p.BadgeLevel,
		)

		// -------------------------------
		// QUÊTE ACTIVE
		// -------------------------------

		if p.ActiveQuestID != "" {

			activeQuest := getActiveQuest(p)

			if activeQuest != nil {

				fmt.Println("\n----- QUÊTE ACTIVE -----")

				showQuestDetails(
					*activeQuest,
					p.QuestProgress,
				)
			}
		}

		// -------------------------------
		// QUÊTES DISPONIBLES
		// -------------------------------

		available := availableQuests(p)

		fmt.Println("\n===================================")
		fmt.Println("         QUÊTES DISPONIBLES")
		fmt.Println("===================================")

		if len(available) == 0 {

			fmt.Println("Aucune nouvelle quête disponible.")

		} else {

			for i, quest := range available {

				fmt.Printf(
					"%d. %s",
					i+1,
					quest.Name,
				)

				if quest.RewardBadge > 0 {
					fmt.Printf(
						" [Badge %d]",
						quest.RewardBadge,
					)
				}

				fmt.Println()
			}
		}

		// -------------------------------
		// OPTIONS
		// -------------------------------

		nextChoice := len(available) + 1

		if p.ActiveQuestID != "" {

			fmt.Printf(
				"%d. Abandonner la quête active\n",
				nextChoice,
			)

			nextChoice++
		}

		fmt.Printf(
			"%d. Retour\n",
			nextChoice,
		)

		fmt.Print("\nEnter your choice: ")

		choice := readIntInput()

		if choice == -1 {

			fmt.Println(
				"\nErreur : Veuillez entrer un nombre valide !",
			)

			continue
		}

		// -------------------------------
		// RETOUR
		// -------------------------------

		if choice == nextChoice {
			return
		}

		// -------------------------------
		// ABANDON
		// -------------------------------

		if p.ActiveQuestID != "" &&
			choice == len(available)+1 {

			fmt.Println(
				"\nAttention : abandonner la quête supprimera votre progression.",
			)

			fmt.Println("1. Confirmer")
			fmt.Println("2. Annuler")

			fmt.Print("Votre choix : ")

			confirm := readIntInput()

			if confirm == 1 {
				abandonQuest(p)
			}

			continue
		}

		// -------------------------------
		// SÉLECTION D'UNE QUÊTE
		// -------------------------------

		if choice >= 1 &&
			choice <= len(available) {

			selectedQuest := available[choice-1]

			fmt.Println()

			showQuestDetails(
				selectedQuest,
				0,
			)

			fmt.Println("\n1. Accepter la quête")
			fmt.Println("2. Retour")

			fmt.Print("Votre choix : ")

			accept := readIntInput()

			if accept != 1 {
				continue
			}

			// -------------------------------
			// CHANGEMENT DE QUÊTE
			// -------------------------------

			if p.ActiveQuestID != "" {

				if selectedQuest.StartItem != "" &&
					isInventoryFull(p) {

					fmt.Println(
						"\nInventaire plein. Impossible d'accepter cette quête.",
					)

					continue
				}

				fmt.Println(
					"\n===================================",
				)

				fmt.Println(
					"Vous avez déjà une quête active.",
				)

				fmt.Println(
					"Changer de quête supprimera sa progression.",
				)

				fmt.Println(
					"1. Changer de quête",
				)

				fmt.Println(
					"2. Garder la quête actuelle",
				)

				fmt.Print("Votre choix : ")

				change := readIntInput()

				if change != 1 {
					continue
				}

				abandonQuest(p)
			}

			acceptQuest(
				p,
				selectedQuest,
			)
		}
	}
}