package main

import "fmt"

type NPC struct {
	Name        string
	Description string
	Dialogue    string
}

func interactWithNPC(p *Character, npc NPC, locationName string) {
	fmt.Println("\n===================================")
	fmt.Println("             PNJ")
	fmt.Println("===================================")
	fmt.Println(npc.Name)
	fmt.Println(npc.Description)
	fmt.Println("\n\"" + npc.Dialogue + "\"")
	quest := getActiveQuest(p)
	if quest == nil {
		return
	}
	if quest.ObjectiveType != "delivery" {
		return
	}
	if quest.TargetNPC != npc.Name {
		return
	}
	if quest.TargetLocation != locationName {
		return
	}
	if quest.StartItem != "" {
		if !hasInventoryItem(p, quest.StartItem) {
			fmt.Println("\nLe PNJ attend un colis, mais vous ne l'avez plus.")
			return
		}
		removeInventoryItem(p, quest.StartItem)
	}
	fmt.Println("\nLe colis a été remis avec succès !")
	updateQuestProgress(
		p,
		"delivery",
		quest.TargetNPC,
	)
}
