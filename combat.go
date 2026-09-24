package main

import "fmt"

func MonstrePattern(
	p *Character,
	monster *Monster,
	turn int,
) {
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

func characterTurn(
	p *Character,
	monster *Monster,
) bool {
	for {
		printPanel("Tour de " + p.Name)
		printOption(1, "Attaque basique")
		printOption(2, "Competences cybernetiques")
		printOption(3, "Inventaire")
		var choice int
		askChoice()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			damage := p.Attack
			monster.HP -= damage
			if monster.HP < 0 {
				monster.HP = 0
			}
			fmt.Println("Attaque utilisee : Attaque basique")
			fmt.Println(
				"Degats infliges :",
				damage,
			)
			fmt.Printf(
				"%s : %d/%d HP\n",
				monster.Name,
				monster.HP,
				monster.MaxHP,
			)
			return true
		case 2:
			played := combatSkills(
				p,
				monster,
			)
			if played {
				return true
			}
			continue
		case 3:
			played := combatInventory(
				p,
				monster,
			)
			if played {
				return true
			}
			continue
		default:
			fmt.Println(
				"Choix invalide.",
			)
		}
	}
}

func combatSkills(
	p *Character,
	monster *Monster,
) bool {
	printPanel("Competences cybernetiques")
	if len(p.Skill) == 0 {
		fmt.Println(
			"Aucune compétence disponible.",
		)
		return false
	}
	for i, skill := range p.Skill {
		fmt.Printf(
			"%d. %s\n",
			i+1,
			skill,
		)
	}
	fmt.Println("0. Retour")
	var choice int
	askChoice()
	fmt.Scanln(&choice)
	if choice == 0 {
		return false
	}
	if choice < 1 ||
		choice > len(p.Skill) {
		fmt.Println(
			"Choix invalide.",
		)
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
		fmt.Println(
			"Cette competence n'est pas utilisable.",
		)
		return false
	}
	if p.Energy < energyCost {
		fmt.Println(
			"\nEnergie insuffisante.",
		)
		fmt.Printf(
			"Energie necessaire : %d | Energie actuelle : %d\n",
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
	fmt.Printf(
		"\n%s utilisee.\n",
		skill,
	)
	fmt.Printf(
		"Degats infliges : %d\n",
		damage,
	)
	fmt.Printf(
		"Energie consommee : %d\n",
		energyCost,
	)
	fmt.Printf(
		"Energie restante : %d/%d\n",
		p.Energy,
		p.MaxEnergy,
	)
	fmt.Printf(
		"%s : %d/%d HP\n",
		monster.Name,
		monster.HP,
		monster.MaxHP,
	)
	return true
}

// ZONE D'ENTRAÎNEMENT
func trainingFight(p *Character) {
	monster := initRobot()
	fight(p, monster)
}

// MONSTRE VAINCU
func monsterDefeated(
	p *Character,
	monster *Monster,
) {
	updateQuestProgress(
		p,
		"kill",
		monster.Name,
	)
	giveDrops(
		p,
		monster,
	)
	gainExperience(
		p,
		monster.XPReward,
	)
}

// COMBAT GÉNÉRAL
func fight(
	p *Character,
	monster Monster,
) {
	turn := 1
	for p.HP > 0 &&
		monster.HP > 0 {
		printHeader(fmt.Sprintf("Tour %d", turn))
		// Joueur commence
		if p.Initiative >= monster.Initiative {
			fmt.Println(
				">>",
				p.Name,
				"commence le combat !",
			)
			playerPlayed := characterTurn(
				p,
				&monster,
			)
			if !playerPlayed {
				continue
			}
			if monster.HP <= 0 {
				monsterDefeated(
					p,
					&monster,
				)
				break
			}
			poisonTick(&monster)
			if monster.HP <= 0 {
				monsterDefeated(
					p,
					&monster,
				)
				break
			}
			MonstrePattern(
				p,
				&monster,
				turn,
			)
			if isDead(p) {
				break
			}
		}
		// Monstre commence
		if p.Initiative < monster.Initiative {
			fmt.Println(
				">>",
				monster.Name,
				"commence le combat !",
			)
			MonstrePattern(
				p,
				&monster,
				turn,
			)
			if isDead(p) {
				break
			}
			poisonTick(&monster)
			if monster.HP <= 0 {
				monsterDefeated(
					p,
					&monster,
				)
				break
			}
			playerPlayed := characterTurn(
				p,
				&monster,
			)
			if !playerPlayed {
				continue
			}
			if monster.HP <= 0 {
				monsterDefeated(
					p,
					&monster,
				)
				break
			}
		}
		turn++
	}
}
