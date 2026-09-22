package main

import "fmt"

func goblinPattern(p *Character, monster *Monster, turn int) {
	damage := monster.Attack
	// Tous les 3 tours, le gobelin inflige 200% de son attaque
	if turn%3 == 0 {
		damage = monster.Attack * 2
	}
	p.HP -= damage
	fmt.Printf(
		"%s inflige à %s %d de dégâts\n",
		monster.Name,
		p.Name,
		damage,
	)
	fmt.Printf(
		"%s : %d/%d HP\n",
		p.Name,
		p.HP,
		p.MaxHP,
	)
}

func characterTurn(p *Character, monster *Monster) bool {
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		monster.HP -= 5
		fmt.Println("Attaque utilisée : Attaque basique")
		fmt.Println("Dégâts infligés :", 5)
		fmt.Printf(
			"%s : %d/%d HP\n",
			monster.Name,
			monster.HP,
			monster.MaxHP,
		)
		return true
	case 2:
		return combatInventory(p, monster)
	default:
		fmt.Println("Choix invalide.")
		return false
	}
}

func trainingFight(p *Character) {
	monster := initGoblin()
	turn := 1
	for p.HP > 0 && monster.HP > 0 {
		fmt.Println("\n===================================")
		fmt.Println("             TOUR", turn)
		fmt.Println("===================================")
		playerPlayed := characterTurn(p, &monster)
		if !playerPlayed {
			continue
		}
		if monster.HP <= 0 {
			break
		}
		poisonTick(&monster)
		if monster.HP <= 0 {
			break
		}
		goblinPattern(p, &monster, turn)
		if isDead(p) {
			break
		}
		turn++
	}
}
