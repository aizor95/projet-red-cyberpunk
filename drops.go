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
			if addInventory(p, drop.ItemName) {
				fmt.Printf("Drop obtenu : %s !\n", drop.ItemName)
				dropObtained = true
			} else {
				fmt.Printf(
					"Le drop %s n'a pas pu être récupéré : inventaire plein.\n",
					drop.ItemName,
				)
			}
		}
	}
	if !dropObtained {
		fmt.Println("Aucun matériau obtenu.")
	}
	fmt.Println("============================")
}
