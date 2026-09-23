package main

import "fmt"

func MonstrePattern(p *Character, monster *Monster, turn int) {
	damage := monster.Attack
	// Tous les 3 tours, le monstre inflige 200% de son attaque
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
	for {
		fmt.Println("\n===== TOUR DE", p.Name, "=====")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Compétences cybernétiques")
		fmt.Println("3. Inventaire")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			monster.HP -= 5
			if monster.HP < 0 {
				monster.HP = 0
			}
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
    		played := combatSkills(p, monster)
			if played {
				return true
			}
			continue
		case 3:
			played := combatInventory(p, monster)
			if played {
				return true
			}
			continue
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func combatSkills(p *Character, monster *Monster) bool {
	fmt.Println("\n===== COMPÉTENCES CYBERNÉTIQUES =====")
	if len(p.Skill) == 0 {
		fmt.Println("Aucune compétence disponible.")
		return false
	}
	for i, skill := range p.Skill {
		fmt.Printf("%d. %s\n", i+1, skill)
	}
	fmt.Println("0. Retour")
	var choice int
	fmt.Print("Choisissez une compétence : ")
	fmt.Scanln(&choice)
	if choice == 0 {
		return false
	}
	if choice < 1 || choice > len(p.Skill) {
		fmt.Println("Choix invalide.")
		return false
	}
	skill := p.Skill[choice-1]
	damage := 0
	energyCost := 0
	switch skill {
	case "Frappe cybernétique":
		damage = 8
		energyCost = 10
	case "Surcharge plasma":
		damage = 18
		energyCost = 25
	default:
		fmt.Println("Cette compétence n'est pas utilisable.")
		return false
	}
	if p.Energy < energyCost {
		fmt.Println("\n[!] Energy insuffisante !")
		fmt.Printf(
			"Energy nécessaire : %d | Energy actuelle : %d\n",
			energyCost,
			p.Energy,
		)
		return false
	}
	p.Energy -= energyCost
	monster.HP -= damage
	if monster.HP < 0 {
		monster.HP = 0
	}
	fmt.Printf("\n>>> %s utilisé !\n", skill)
	fmt.Printf("Dégâts infligés : %d\n", damage)
	fmt.Printf("Energy consommée : %d\n",energyCost,)
	fmt.Printf("Energy restante : %d/%d\n",p.Energy,p.MaxEnergy,)
	fmt.Printf("%s : %d/%d HP\n", monster.Name, monster.HP, monster.MaxHP)
	return true
}

func trainingFight(p *Character) {
	monster := initRobot()
	fight(p, monster)
}

func monsterDefeated(p *Character, monster *Monster) {
	if missionKillActive {
		countKills++
		fmt.Printf(
			"\nMission Kill 3 Monsters : %d/3\n",
			countKills,
		)
	}
	giveDrops(p, monster)
	gainExperience(p, monster.XPReward)
}

func fight(p *Character, monster Monster) {
	turn := 1
	for p.HP > 0 && monster.HP > 0 {
		fmt.Println("\n===================================")
		fmt.Println("             TOUR", turn)
		fmt.Println("===================================")
		if p.Initiative >= monster.Initiative {
			fmt.Println(">>", p.Name, "commence le combat !")
			playerPlayed := characterTurn(p, &monster)
			if !playerPlayed {
				continue
			}
			if monster.HP <= 0 {
				monsterDefeated(p, &monster)
				break
			}
			poisonTick(&monster)
			if monster.HP <= 0 {
				monsterDefeated(p, &monster)
				break
			}
			MonstrePattern(p, &monster, turn)
			if isDead(p) {
				break
			}
		}
		if p.Initiative < monster.Initiative {
			fmt.Println(">>", monster.Name, "commence le combat !")
			MonstrePattern(p, &monster, turn)
			if isDead(p) {
				break
			}
			poisonTick(&monster)
			if monster.HP <= 0 {
				monsterDefeated(p, &monster)
				break
			}
			playerPlayed := characterTurn(p, &monster)
			if !playerPlayed {
				continue
			}
			if monster.HP <= 0 {
				monsterDefeated(p, &monster)
				break
			}
		}
		turn++
	}
}
