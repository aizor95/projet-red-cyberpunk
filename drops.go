package main

import (
	"fmt"
	"math/rand"
)

func giveDrops(p *Character, monster *Monster) {
	fmt.Println("\n========== DROPS ==========")
	dropObtained := false
	for _, drop := range monster.Drops {
		random := rand.Intn(100) + 1
		if random <= drop.Chance {
			addInventory(p, drop.ItemName)
			fmt.Printf("Drop obtenu : %s !\n", drop.ItemName)
			dropObtained = true
		}
	}
	if !dropObtained {
		fmt.Println("Aucun matériau obtenu.")
	}
	fmt.Println("============================")
}