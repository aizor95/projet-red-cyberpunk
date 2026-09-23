package main

import "fmt"

func EquipeArmor(p *Character) {
	fmt.Println("\n===================================")
	fmt.Println("         EQUIPMENT MANAGEMENT      ")
	fmt.Println("===================================")

	if len(p.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}

	fmt.Println("Choose an item to equip:")
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Back")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)

	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(p.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}

	itemName := p.Inventory[choice-1]

	var slot string
	var hpBonus int

	switch itemName {
	case "Visière Tactique":
		slot = "head"
		hpBonus = 10
	case "Veste en Fibre de Carbone":
		slot = "torse"
		hpBonus = 25
	case "Bottes Cybernétiques":
		slot = "feet"
		hpBonus = 15
	default:
		fmt.Println("This item cannot be equipped.")
		return
	}

	p.Inventory = append(p.Inventory[:choice-1], p.Inventory[choice:]...)

	switch slot {
	case "head":
		if p.Equipment.head.Name != "" {
			p.Inventory = append(p.Inventory, p.Equipment.head.Name)
			p.MaxHP -= p.Equipment.head.MaxHPbonus
		}
		p.Equipment.head = EquipmentItem{Name: itemName, Slot: "head", MaxHPbonus: hpBonus}
	case "torse":
		if p.Equipment.torse.Name != "" {
			p.Inventory = append(p.Inventory, p.Equipment.torse.Name)
			p.MaxHP -= p.Equipment.torse.MaxHPbonus
		}
		p.Equipment.torse = EquipmentItem{Name: itemName, Slot: "torse", MaxHPbonus: hpBonus}
	case "feet":
		if p.Equipment.feet.Name != "" {
			p.Inventory = append(p.Inventory, p.Equipment.feet.Name)
			p.MaxHP -= p.Equipment.feet.MaxHPbonus
		}
		p.Equipment.feet = EquipmentItem{Name: itemName, Slot: "feet", MaxHPbonus: hpBonus}
	}

	p.MaxHP += hpBonus
	p.HP += hpBonus

	fmt.Printf("\nSuccess! You equipped: %s (+%d max HP)\n", itemName, hpBonus)
}