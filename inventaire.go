package main

import (
	"fmt"
	"time"
)

// INVENTAIRE

func accessInventory(p *Character) {
	fmt.Println("\n===================================")
	fmt.Println("          STORAGE UNIT")
	fmt.Println("===================================")
	if len(p.Inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Printf("\nStorage: %d/%d items\n", len(p.Inventory), p.MaxInventory)
}

// Vérifie si l'inventaire est plein
func isInventoryFull(p *Character) bool {
	return len(p.Inventory) >= p.MaxInventory
}

// Ajoute un objet
func addInventory(p *Character, item string) {
	if isInventoryFull(p) {
		fmt.Println("\n[!] WARNING: Storage capacity reached!")
		fmt.Printf("Limit is %d items.\n", p.MaxInventory)
		fmt.Println("Could not acquire:", item)
		return
	}
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("\n>> Added to storage: %s\n", item)
	fmt.Printf("Storage: %d/%d items\n", len(p.Inventory), p.MaxInventory)
}

// Retire le premier exemplaire d'un objet
func removeInventoryItem(p *Character, itemName string) bool {
	for i, item := range p.Inventory {
		if item == itemName {
			// On supprime l'objet situé à l'indice i
			p.Inventory = append(
				p.Inventory[:i],
				p.Inventory[i+1:]...,
			)
			return true
		}
	}
	return false
}

// UTILISATION DES OBJETS

// Stimpack
func takePot(p *Character) {
	if !removeInventoryItem(p, "Stimpack") {
		fmt.Println("\nError: No Stimpack available!")
		return
	}
	p.HP += 50
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}
	fmt.Println("\n>>> Stimpack used! (+50 HP)")
	fmt.Printf("Current body integrity: %d/%d HP\n", p.HP, p.MaxHP)
}

// Cyber Virus
func poisonPot(p *Character) {
	if !removeInventoryItem(p, "Cyber Virus") {
		fmt.Println("\nError: No Cyber Virus available!")
		return
	}
	fmt.Println("\n[!] Cyber Virus activated!")
	fmt.Println("You are taking damage over time...")
	for tick := 1; tick <= 3; tick++ {
		time.Sleep(1 * time.Second)
		p.HP -= 10
		fmt.Printf(
			"Tick %d/3 - HP: %d/%d\n",
			tick,
			p.HP,
			p.MaxHP,
		)

		if isDead(p) {
			break
		}
	}
}

// Livre de sorts
func spellBook(p *Character) {
	for _, skill := range p.Skill {
		if skill == "Surcharge plasma" {
			fmt.Println("\nYou have already learned Surcharge plasma!")
			return
		}
	}
	p.Skill = append(p.Skill, "Surcharge plasma")
	fmt.Println("\n>>> New programme learned: Surcharge plasma!")
}

// UTILISATION DE L'INVENTAIRE

func inventoryMenu(p *Character) {
	for {
		accessInventory(p)
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Use Stimpack")
		fmt.Println("2. Use Cyber Virus")
		fmt.Println("3. Use Programme : Surcharge plasma")
		fmt.Println("4. Back")
		var choice int
		fmt.Print("Choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			takePot(p)
		case 2:
			poisonPot(p)
		case 3:
			if removeInventoryItem(p, "Programme : Surcharge plasma") {
				spellBook(p)
			} else {
				fmt.Println("\nYou do not have this programme.")
			}
		case 4:
			return
		default:
			fmt.Println("\nInvalid choice.")
		}
	}
}

func combatInventory(p *Character, monster *Monster) bool {
	fmt.Println("\n===== INVENTAIRE DE COMBAT =====")
	if len(p.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return false
	}
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")
	var choice int
	fmt.Print("Choisissez un objet : ")
	fmt.Scanln(&choice)
	if choice == 0 {
		return false
	}
	if choice < 1 || choice > len(p.Inventory) {
		fmt.Println("Choix invalide.")
		return false
	}
	item := p.Inventory[choice-1]
	switch item {
	case "Stimpack":
		takePot(p)
	case "Cyber Virus":
		removeInventoryItem(p, "Cyber Virus")
		poisonMonster(monster)
	case "Programme : Surcharge plasma":
		removeInventoryItem(p, "Programme : Surcharge plasma")
		spellBook(p)
	default:
		fmt.Println("Cet objet ne peut pas être utilisé.")
		return false
	}
	return true
}

func poisonMonster(monster *Monster) {
	monster.Poisoned = true
	monster.PoisonTurns = 3
	fmt.Println("\n[!] Cyber Virus activé sur", monster.Name)
	fmt.Println("L'ennemi sera empoisonné pendant 3 tours.")
}

func poisonTick(monster *Monster) {
	if !monster.Poisoned {
		return
	}
	monster.HP -= 10
	monster.PoisonTurns--
	if monster.HP < 0 {
		monster.HP = 0
	}
	fmt.Printf(
		"\n[POISON] %s perd 10 HP : %d/%d HP\n",
		monster.Name,
		monster.HP,
		monster.MaxHP,
	)
	if monster.PoisonTurns <= 0 {
		monster.Poisoned = false
		fmt.Println("[POISON] Le Cyber Virus n'agit plus.")
	}
}

// amélioration de l'inventaire

func upgradeInventorySlot(p *Character) {
	if p.UpgradeCount >= 3 {
		fmt.Println("Maximum inventory upgrades reached! (Max 3 times)")
		return
	}
	cost := 30
	if p.Argent < cost {
		fmt.Println("Not enough credits")
		return
	}
	p.Argent -= cost
	p.MaxInventory += 10 
	p.UpgradeCount++     
	fmt.Println("Inventory upgraded!")
	fmt.Printf("New capacity: %d\n", p.MaxInventory)
}