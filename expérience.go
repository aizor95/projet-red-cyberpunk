package main

import "fmt"

func gainExperience(p *Character, amount int) {
	p.Experience += amount
	fmt.Printf("\n+%d XP de réputation\n", amount)
	fmt.Printf("XP : %d/%d\n", p.Experience, p.MaxExperience)
	for p.Experience >= p.MaxExperience {
		p.Experience -= p.MaxExperience
		levelUp(p)
	}
}

func levelUp(p *Character) {
	p.Level++
	p.MaxExperience += 25
	p.MaxHP += 10
	p.HP = p.MaxHP
	p.Attack += 2
	p.Defense += 1
	p.MaxEnergy += 10
	p.Energy += 10
	if p.Energy > p.MaxEnergy {
		p.Energy = p.MaxEnergy
	}
	fmt.Println("\n===================================")
	fmt.Println("       NIVEAU DE RÉPUTATION +1")
	fmt.Println("===================================")
	fmt.Println("Niveau :", p.Level)
	fmt.Println("Max HP :", p.MaxHP)
	fmt.Println("Attack :", p.Attack)
	fmt.Println("Defense :", p.Defense)
	fmt.Println("Energy :", p.Energy)
	fmt.Println("===================================")
}