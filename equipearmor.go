package main

import "fmt"

func EquipeArmor(p *Character) {
	printHeader("Equipement")
	fmt.Println("Installations actuelles :")
	if p.Equipment.head.Name != "" {
		fmt.Printf("- Tete   : %s (+%d HP max)\n", p.Equipment.head.Name, p.Equipment.head.MaxHPbonus)
	} else {
		fmt.Println("- Tete   : vide")
	}
	if p.Equipment.torse.Name != "" {
		fmt.Printf("- Torse  : %s (+%d HP max)\n", p.Equipment.torse.Name, p.Equipment.torse.MaxHPbonus)
	} else {
		fmt.Println("- Torse  : vide")
	}
	if p.Equipment.feet.Name != "" {
		fmt.Printf("- Pieds  : %s (+%d HP max)\n", p.Equipment.feet.Name, p.Equipment.feet.MaxHPbonus)
	} else {
		fmt.Println("- Pieds  : vide")
	}
	var armorItems []string
	for _, item := range p.Inventory {
		if item == "Visière tactique" || item == "Veste en fibre de carbone" || item == "Bottes cybernétiques" {
			armorItems = append(armorItems, item)
		}
	}
	if len(armorItems) == 0 {
		fmt.Println("\nAucun equipement a équiper dans l'inventaire.")
		return
	}
	fmt.Println("\nEquipement disponible :")
	for i, item := range armorItems {
		printOption(i+1, item)
	}
	fmt.Println("  0. Retour")
	var choice int
	askChoice()
	fmt.Scanln(&choice)
	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(armorItems) {
		fmt.Println("Choix invalide.")
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
	if !removeInventoryItem(p, itemName) {
		fmt.Println("Impossible de retirer l'équipement de l'inventaire.")
		return
	}
	var oldEquipment EquipmentItem
	switch slot {
	case "head":
		oldEquipment = p.Equipment.head
	case "torse":
		oldEquipment = p.Equipment.torse
	case "feet":
		oldEquipment = p.Equipment.feet
	}
	oldBonus := oldEquipment.MaxHPbonus
	if oldEquipment.Name != "" {
		p.Inventory = append(p.Inventory, oldEquipment.Name)
	}
	p.MaxHP += hpBonus - oldBonus
	p.HP += hpBonus - oldBonus
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}
	if p.HP < 0 {
		p.HP = 0
	}
	switch slot {
	case "head":
		p.Equipment.head = EquipmentItem{
			Name:       itemName,
			Slot:       "head",
			MaxHPbonus: hpBonus,
		}
	case "torse":
		p.Equipment.torse = EquipmentItem{
			Name:       itemName,
			Slot:       "torse",
			MaxHPbonus: hpBonus,
		}
	case "feet":
		p.Equipment.feet = EquipmentItem{
			Name:       itemName,
			Slot:       "feet",
			MaxHPbonus: hpBonus,
		}
	}
	fmt.Printf("\nEquipement installe : %s (+%d HP max)\n", itemName, hpBonus)
	fmt.Printf("Vie actuelle : %d/%d\n", p.HP, p.MaxHP)
}
