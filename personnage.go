package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Character struct {
	Name          string
	Class         string
	Level         int
	HP            int
	MaxHP         int
	Energy        int
	MaxEnergy     int
	Attack        int
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
	Energy     int
}

func initCharacter(name string, class string) Character {
	maxHP := 0
	energy := 0
	attack := 0
	initiative := 0
	switch class {
	case "Netrunner":
		maxHP = 80
		energy = 120
		attack = 5
		initiative = 10
	case "Mercenary":
		maxHP = 100
		energy = 100
		attack = 7
		initiative = 7
	case "Cyborg":
		maxHP = 120
		energy = 80
		attack = 6
		initiative = 5
	}
	return Character{
		Name:            name,
		Class:           class,
		Level:           1,
		HP:              maxHP / 2,
		MaxHP:           maxHP,
		Energy:          energy,
		MaxEnergy:       energy,
		Attack:          attack,
		Initiative:      initiative,
		Experience:      0,
		MaxExperience:   100,
		Argent:          100,
		Equipment:       Equipment{},
		Inventory:       []string{"Stimpack", "Cyber Virus"},
		Skill:           []string{"Frappe cybernétique"},
		MaxInventory:    10,
		UpgradeCount:    0,
		CurrentLocation: "Place centrale",
		BadgeLevel:      0,
		ActiveQuestID:   "",
		QuestProgress:   0,
		CompletedQuests: []string{},
	}
}

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

func characterCreation() Character {
	var rawName string
	var choice int
	var selectedClass string
	for {
		fmt.Print("\nNom du personnage (lettres uniquement) : ")
		fmt.Scanln(&rawName)
		if isValidName(rawName) {
			break
		}
		fmt.Println("Nom invalide. Utilise seulement des lettres.")
	}
	formattedName := formatName(rawName)
	printPanel("Choix de la classe")
	printOption(1, "Netrunner  - rapide, beaucoup d'energie")
	printOption(2, "Mercenary  - equilibre, bonne attaque")
	printOption(3, "Cyborg     - solide, plus de vie")
	for {
		askChoice()
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Choix invalide. Entre un nombre entre 1 et 3.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if choice >= 1 && choice <= 3 {
			break
		}
		fmt.Println("Choix invalide. Entre un nombre entre 1 et 3.")
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

func displayInfo(p *Character) {
	printHeader("Profil")
	fmt.Println("Nom        :", p.Name)
	fmt.Println("Classe     :", p.Class)
	fmt.Println("Niveau     :", p.Level)
	fmt.Printf(
		"Vie        : %d/%d HP\n",
		p.HP,
		p.MaxHP,
	)
	fmt.Printf(
		"Energie    : %d/%d\n",
		p.Energy,
		p.MaxEnergy,
	)
	fmt.Println("Attaque    :", p.Attack)
	fmt.Println("Credits    :", p.Argent)
	fmt.Println(
		"Badge      : niveau",
		p.BadgeLevel,
	)
	fmt.Println(
		"Position   :",
		p.CurrentLocation,
	)
	if p.ActiveQuestID != "" {
		quest := getActiveQuest(p)
		if quest != nil {
			fmt.Printf(
				"Mission    : %s (%d/%d)\n",
				quest.Name,
				p.QuestProgress,
				quest.RequiredAmount,
			)
		}
	}
	fmt.Println("Competences:", p.Skill)
}

func isDead(p *Character) bool {
	if p.HP <= 0 {
		printHeader("System failure")
		fmt.Println("Tu t'effondres dans les rues de Neon City...")
		p.HP = p.MaxHP / 2
		fmt.Println("\nRedemarrage des implants...")
		fmt.Printf(
			"Tu reviens avec %d/%d HP.\n",
			p.HP,
			p.MaxHP,
		)
		return true
	}
	return false
}
