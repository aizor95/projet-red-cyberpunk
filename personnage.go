package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ===================================
// STRUCTURES
// ===================================

type Character struct {
	Name          string
	Class         string
	Level         int
	HP            int
	MaxHP         int
	Energy        int
	MaxEnergy     int
	Attack        int
	Defense       int
	Initiative    int
	Experience    int
	MaxExperience int
	Argent        int

	Equipment Equipment
	Inventory []string
	Skill     []string

	MaxInventory int
	UpgradeCount int

	CurrentLocation string
	BadgeLevel      int

	ActiveQuestID   string
	QuestProgress   int
	CompletedQuests []string
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

// ===================================
// CRÉATION DU PERSONNAGE
// ===================================

func initCharacter(name string, class string) Character {

	maxHP := 0
	energy := 0
	attack := 0
	defense := 0
	initiative := 0

	switch class {

	case "Netrunner":

		maxHP = 80
		energy = 120
		attack = 5
		defense = 2
		initiative = 10

	case "Mercenary":

		maxHP = 100
		energy = 100
		attack = 7
		defense = 5
		initiative = 7

	case "Cyborg":

		maxHP = 120
		energy = 80
		attack = 6
		defense = 8
		initiative = 5
	}

	return Character{
		Name:          name,
		Class:         class,
		Level:         1,
		HP:            maxHP / 2,
		MaxHP:         maxHP,
		Energy:        energy,
		MaxEnergy:     energy,
		Attack:        attack,
		Defense:       defense,
		Initiative:    initiative,
		Experience:    0,
		MaxExperience: 100,
		Argent:        100,
		Equipment:     Equipment{},
		Inventory:     []string{"Stimpack", "Cyber Virus"},
		Skill:         []string{"Frappe cybernétique"},
		MaxInventory:  10,
		UpgradeCount:  0,
		CurrentLocation: "Place centrale",
		BadgeLevel:      0,
		ActiveQuestID:   "",
		QuestProgress:   0,
		CompletedQuests: []string{},
	}
}

// ===================================
// NOM DU PERSONNAGE
// ===================================

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

// ===================================
// CRÉATION INTERACTIVE
// ===================================

func characterCreation() Character {

	var rawName string
	var choice int
	var selectedClass string

	for {

		fmt.Print("Enter character name (letters only): ")
		fmt.Scanln(&rawName)

		if isValidName(rawName) {
			break
		}

		fmt.Println(
			"Invalid input! Name must contain letters only.",
		)
	}

	formattedName := formatName(rawName)

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

		fmt.Println(
			"Invalid choice, select between 1 and 3.",
		)
	}

	switch choice {

	case 1:
		selectedClass = "Netrunner"

	case 2:
		selectedClass = "Mercenary"

	case 3:
		selectedClass = "Cyborg"
	}

	return initCharacter(
		formattedName,
		selectedClass,
	)
}

// ===================================
// PROFIL
// ===================================

func displayInfo(p *Character) {

	fmt.Println("\n===================================")
	fmt.Println("         CHARACTER PROFILE")
	fmt.Println("===================================")

	fmt.Println("Name:", p.Name)
	fmt.Println("Class:", p.Class)
	fmt.Println("Level:", p.Level)

	fmt.Printf(
		"Health: %d/%d HP\n",
		p.HP,
		p.MaxHP,
	)

	fmt.Printf(
		"Energy: %d/%d\n",
		p.Energy,
		p.MaxEnergy,
	)

	fmt.Println("Attack:", p.Attack)
	fmt.Println("Defense:", p.Defense)
	fmt.Println("Credits:", p.Argent)

	fmt.Println(
		"Badge level:",
		p.BadgeLevel,
	)

	fmt.Println(
		"Current location:",
		p.CurrentLocation,
	)

	if p.ActiveQuestID != "" {

		quest := getActiveQuest(p)

		if quest != nil {

			fmt.Printf(
				"Active quest: %s (%d/%d)\n",
				quest.Name,
				p.QuestProgress,
				quest.RequiredAmount,
			)
		}
	}

	fmt.Println("Skills:", p.Skill)
	fmt.Println("===================================")
}

// ===================================
// MORT
// ===================================

func isDead(p *Character) bool {

	if p.HP <= 0 {

		fmt.Println("\n===================================")
		fmt.Println("          SYSTEM FAILURE")
		fmt.Println("              WASTED")
		fmt.Println("===================================")

		fmt.Println(
			"You died in the streets of Neon City...",
		)

		p.HP = p.MaxHP / 2

		fmt.Println(
			"\nSystem reboot in progress...",
		)

		fmt.Printf(
			"You have been revived with %d/%d HP.\n",
			p.HP,
			p.MaxHP,
		)

		fmt.Println("===================================")

		return true
	}

	return false
}