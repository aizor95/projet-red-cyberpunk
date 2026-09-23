package main

import (
	"fmt"
	"math/rand"
)

func randomEncounter(p *Character, location Location) {
	if len(location.Enemies) == 0 {
		fmt.Println("\nAucun ennemi dans cette zone.")
		return
	}
	random := rand.Intn(len(location.Enemies))
	monsterName := location.Enemies[random]
	monster := createMonster(monsterName)
	if monster.Name == "" {
		fmt.Println("\nErreur : impossible de créer le monstre.")
		return
	}
	fmt.Println("\n===================================")
	fmt.Println("            RENCONTRE")
	fmt.Println("===================================")
	fmt.Println("Un ennemi apparaît :", monster.Name)
	fight(p, monster)
}