package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// STRUCTURES

type Character struct {
	Name      string
	Class     string
	Level     int
	HP        int
	MaxHP     int
	Energy    int
	Attack    int
	Defense   int
	Argent    int
	Equipment Equipment
	Inventory []string
	Skill     []string
}

type Equipment struct {
	head  EquipmentItem
	torse EquipmentItem
	feet  EquipmentItem
}

type EquipmentItem struct {
	Name       string
	Slot       string
	MaxHPbonus int
	Attack     int
	Defense    int
	Energy     int
}

// CREATION DU PERSONNAGE

func initCharacter(name string, class string) Character {
	maxHP := 0
	energy := 0
	attack := 0
	defense := 0
	switch class {
	case "Netrunner":
		maxHP = 80
		energy = 120
		attack = 5
		defense = 2
	case "Mercenaire":
		maxHP = 100
		energy = 100
		attack = 7
		defense = 5
	case "Cyborg":
		maxHP = 120
		energy = 80
		attack = 6
		defense = 8
	}

	return Character{
		Name:      name,
		Class:     class,
		Level:     1,
		HP:        maxHP,
		MaxHP:     maxHP,
		Energy:    energy,
		Attack:    attack,
		Defense:   defense,
		Argent:    100,
		Equipment: Equipment{},
		Inventory: []string{"Stimpack", "Cyber Virus"},
		Skill:     []string{"Coup de poing"},
	}
}

// NOM DU PERSONNAGE

func isValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for _, char := range name {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func formatName(name string) string {
	if len(name) == 0 {
		return name
	}
	runes := []rune(strings.ToLower(name))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// CREATION INTERACTIVE

func characterCreation() Character {
	var rawName string
	var choice int
	var selectedClass string
	// Choix du nom
	for {
		fmt.Print("Enter character name (letters only): ")
		fmt.Scanln(&rawName)

		if isValidName(rawName) {
			break
		}
		fmt.Println("Invalid input! Name must contain letters only.")
	}
	formattedName := formatName(rawName)
	// Choix de la classe
	fmt.Println("\nSelect your class:")
	fmt.Println("1. Netrunner")
	fmt.Println("2. Mercenary")
	fmt.Println("3. Cyborg")
	for {
		fmt.Print("Your choice (1-3): ")
		fmt.Scanln(&choice)
		if choice >= 1 && choice <= 3 {
			break
		}
		fmt.Println("Invalid choice, select between 1 and 3.")
	}
	switch choice {
	case 1:
		selectedClass = "Netrunner"
	case 2:
		selectedClass = "Mercenary"
	case 3:
		selectedClass = "Cyborg"
	}
	return initCharacter(formattedName, selectedClass)
}

// AFFICHAGE DU PROFIL

func displayInfo(p *Character) {
	fmt.Println("\n===================================")
	fmt.Println("         CHARACTER PROFILE")
	fmt.Println("===================================")
	fmt.Println("Name:", p.Name)
	fmt.Println("Class:", p.Class)
	fmt.Println("Level:", p.Level)
	fmt.Printf("Health: %d/%d HP\n", p.HP, p.MaxHP)
	fmt.Printf("Energy: %d\n", p.Energy)
	fmt.Println("Attack:", p.Attack)
	fmt.Println("Defense:", p.Defense)
	fmt.Println("Credits:", p.Argent)
	fmt.Println("Skills:", p.Skill)
	fmt.Println("===================================")
}

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
	fmt.Printf("\nStorage: %d/10 items\n", len(p.Inventory))
}

// Vérifie si l'inventaire est plein
func isInventoryFull(p *Character) bool {
	maxCapacity := 10
	if len(p.Inventory) >= maxCapacity {
		return true
	}
	return false
}
func main(){
	var name string
	var choice int
	var class string
	fmt.Println("===========================")
	fmt.Println("NEON CITY")
	fmt.Println("===========================")
	fmt.Println()
	fmt.Println("Création de votre personnage")
	fmt.Println()
	fmt.Println("Quel est votre nom ?")
	fmt.Scanln(&name)
	fmt.Println()
	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Netrunner")
	fmt.Println("2. Mercenaire")
	fmt.Println("3. Cyborg")
	fmt.Println()
	fmt.Println("Votre choix : (Numéro de la classe)")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		class = "Netrunner"
	case 2:
		class = "Mercenaire"
	case 3:
		class = "Cyborg"
	}
	character := initCharacter(name,class)
	fmt.Println()
	fmt.Println("===========================")
	fmt.Println("Personnage créé")
	fmt.Println("===========================")
	fmt.Println("")
	fmt.Println("bienvenue :",character.Name)
	fmt.Println("Attaque :",character.Attack)
}
