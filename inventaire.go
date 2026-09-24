package main

import (
	"fmt"
	"time"
)

func accessInventory(p *Character) {
	printHeader("Inventaire")
	if len(p.Inventory) == 0 {
		fmt.Println("Ton sac est vide.")
		return
	}
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Printf("\nStockage : %d/%d objets\n", len(p.Inventory), p.MaxInventory)
}

func isInventoryFull(p *Character) bool {
	return len(p.Inventory) >= p.MaxInventory
}

func addInventory(p *Character, item string) bool {
	if isInventoryFull(p) {
		fmt.Println("\nInventaire plein.")
		fmt.Printf("Limite actuelle : %d objets.\n", p.MaxInventory)
		fmt.Println("Objet refuse :", item)
		return false
	}
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("\nObjet ajoute : %s\n", item)
	fmt.Printf("Stockage : %d/%d objets\n", len(p.Inventory), p.MaxInventory)
	return true
}

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

func takePot(p *Character) {
	if !removeInventoryItem(p, "Stimpack") {
		fmt.Println("\nAucun Stimpack disponible.")
		return
	}
	p.HP += 50
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}
	fmt.Println("\nStimpack utilise. +50 HP")
	fmt.Printf("Vie actuelle : %d/%d HP\n", p.HP, p.MaxHP)
}

func poisonPot(p *Character) {
	if !removeInventoryItem(p, "Cyber Virus") {
		fmt.Println("\nAucun Cyber Virus disponible.")
		return
	}
	fmt.Println("\nCyber Virus active.")
	fmt.Println("Le virus te ronge pendant 3 secondes...")
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

func spellBook(p *Character) {
	for _, skill := range p.Skill {
		if skill == "Surcharge plasma" {
			fmt.Println("\nSurcharge plasma est deja installee.")
			return
		}
	}
	p.Skill = append(p.Skill, "Surcharge plasma")
	fmt.Println("\nNouveau programme installe : Surcharge plasma.")
}

func inventoryMenu(p *Character) {
	for {
		accessInventory(p)
		fmt.Println()
		printOption(1, "Utiliser un Stimpack")
		printOption(2, "Utiliser un Cyber Virus")
		printOption(3, "Utiliser une Batterie")
		printOption(4, "Installer Programme : Surcharge plasma")
		printOption(5, "Retour")
		var choice int
		askChoice()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			takePot(p)
		case 2:
			poisonPot(p)
		case 3:
			useBattery(p)
		case 4:
			if removeInventoryItem(p, "Programme : Surcharge plasma") {
				spellBook(p)
			} else {
				fmt.Println("\nTu n'as pas ce programme.")
			}
		case 5:
			return
		default:
			fmt.Println("\nChoix invalide.")
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
	case "Batterie":
		useBattery(p)
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
		fmt.Println("Inventaire deja augmente au maximum.")
		return
	}
	cost := 30
	if p.Argent < cost {
		fmt.Println("Credits insuffisants.")
		return
	}
	p.Argent -= cost
	p.MaxInventory += 10
	p.UpgradeCount++
	fmt.Println("Inventaire augmente.")
	fmt.Printf("Nouvelle capacite : %d objets\n", p.MaxInventory)
}

func useBattery(p *Character) {
	if !removeInventoryItem(p, "Batterie") {
		fmt.Println("\nAucune Batterie disponible.")
		return
	}
	p.Energy += 30
	if p.Energy > p.MaxEnergy {
		p.Energy = p.MaxEnergy
	}
	fmt.Println("\nBatterie utilisee. +30 energie")
	fmt.Printf("Energie actuelle : %d/%d\n", p.Energy, p.MaxEnergy)
}
