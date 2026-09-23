package main

import "fmt"

func EquipeArmor(p *Character) {
	fmt.Println("\n===================================")
	fmt.Println("         EQUIPMENT MANAGEMENT      ")
	fmt.Println("===================================")

	// Mostra l'equipaggiamento attualmente indossato
	fmt.Println("Current Equipment:")
	if p.Equipment.head.Name != "" {
		fmt.Printf("- Head: %s (+%d max HP)\n", p.Equipment.head.Name, p.Equipment.head.MaxHPbonus)
	} else {
		fmt.Println("- Head: Empty")
	}

	if p.Equipment.torse.Name != "" {
		fmt.Printf("- Torso: %s (+%d max HP)\n", p.Equipment.torse.Name, p.Equipment.torse.MaxHPbonus)
	} else {
		fmt.Println("- Torso: Empty")
	}

	if p.Equipment.feet.Name != "" {
		fmt.Printf("- Feet: %s (+%d max HP)\n", p.Equipment.feet.Name, p.Equipment.feet.MaxHPbonus)
	} else {
		fmt.Println("- Feet: Empty")
	}
	fmt.Println("-----------------------------------")

	// Filtriamo solo gli oggetti che sono effettivamente delle armature
	var armorItems []string
	for _, item := range p.Inventory {
		if item == "Visière tactique" || item == "Veste en fibre de carbone" || item == "Bottes cybernétiques" {
			armorItems = append(armorItems, item)
		}
	}

	if len(armorItems) == 0 {
		fmt.Println("You don't have any armor in your storage.")
		return
	}

	fmt.Println("Choose an armor to equip:")
	for i, item := range armorItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Back")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)

	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(armorItems) {
		fmt.Println("Invalid choice.")
		return
	}

	itemName := armorItems[choice-1]

	var slot string
	var hpBonus int

	switch itemName {
	case "Visière tactique":
		slot = "head"
		hpBonus = 10
	case "Veste en fibre de carbone":
		slot = "torse"
		hpBonus = 25
	case "Bottes cybernétiques":
		slot = "feet"
		hpBonus = 15
	}

	removeInventoryItem(p, itemName)

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
