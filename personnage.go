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
	return len(p.Inventory) >= 10
}

// Ajoute un objet
func addInventory(p *Character, item string) {
	if isInventoryFull(p) {
		fmt.Println("\n[!] WARNING: Storage capacity reached!")
		fmt.Println("Limit is 10 items.")
		fmt.Println("Could not acquire:", item)
		return
	}
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("\n>> Added to storage: %s\n", item)
	fmt.Printf("Storage: %d/10 items\n", len(p.Inventory))
}

// Retire le premier exemplaire d'un objet
func removeInventoryItem(p *Character, itemName string) bool {
	for i, item := range p.Inventory {
		if item == itemName {
			// On supprime l'objet situé à l'indice i
			p.Inventory = append(
				p.Inventory[:i],
				p.Inventory[i+1:]...,
			)
			return true
		}
	}
	return false
}

// UTILISATION DES OBJETS

// Stimpack
func takePot(p *Character) {
	if !removeInventoryItem(p, "Stimpack") {
		fmt.Println("\nError: No Stimpack available!")
		return
	}
	p.HP += 50
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}
	fmt.Println("\n>>> Stimpack used! (+50 HP)")
	fmt.Printf("Current body integrity: %d/%d HP\n", p.HP, p.MaxHP)
}

// Cyber Virus
func poisonPot(p *Character) {
	if !removeInventoryItem(p, "Cyber Virus") {
		fmt.Println("\nError: No Cyber Virus available!")
		return
	}
	fmt.Println("\n[!] Cyber Virus activated!")
	fmt.Println("You are taking damage over time...")
	for tick := 1; tick <= 3; tick++ {
		time.Sleep(1 * time.Second)
		p.HP -= 10
		fmt.Printf(
			"Tick %d/3 - HP: %d/%d\n",
			tick,
			p.HP,
			p.MaxHP,
		)

		if isDead(p) {
			break
		}
	}
}

// Livre de sorts
func spellBook(p *Character) {
	for _, skill := range p.Skill {
		if skill == "Boule de Feu" {
			fmt.Println("\nYou have already learned Boule de Feu!")
			return
		}
	}
	p.Skill = append(p.Skill, "Boule de Feu")
	fmt.Println("\n>>> New spell learned: Boule de Feu!")
}

// UTILISATION DE L'INVENTAIRE

func inventoryMenu(p *Character) {
	for {
		accessInventory(p)
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Use Stimpack")
		fmt.Println("2. Use Cyber Virus")
		fmt.Println("3. Use Spellbook: Boule de Feu")
		fmt.Println("4. Back")
		var choice int
		fmt.Print("Choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			takePot(p)
		case 2:
			poisonPot(p)
		case 3:
			if removeInventoryItem(p, "Spellbook: Boule de Feu") {
				spellBook(p)
			} else {
				fmt.Println("\nYou do not have this spellbook.")
			}
		case 4:
			return
		default:
			fmt.Println("\nInvalid choice.")
		}
	}
}

// MARCHAND

func merchant(p *Character) {
	for {
		fmt.Println("\n===================================")
		fmt.Println("       CYBERNETIC WORKSHOP")
		fmt.Println("===================================")
		fmt.Println("Your credits:", p.Argent)
		fmt.Println()
		fmt.Println("1. Stimpack (3 credits)")
		fmt.Println("2. Cyber Virus (6 credits)")
		fmt.Println("3. Spellbook: Boule de Feu (25 credits)")
		fmt.Println("4. Fourrure de Loup (4 credits)")
		fmt.Println("5. Peau de Troll (7 credits)")
		fmt.Println("6. Cuir de Sanglier (3 credits)")
		fmt.Println("7. Plume de Corbeau (1 credits)")
		fmt.Println("8. Back")

		var choice int
		fmt.Print("\nYour choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if len(p.Inventory) >= 10 {
				fmt.Println("Inventory is full")
			} else if p.Argent < 3 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 3
				addInventory(p, "Stimpack")
				fmt.Println("You bought Stimpack")
			}

		case 2:
			if len(p.Inventory) >= 10 {
				fmt.Println("Inventory is full")
			} else if p.Argent < 6 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 6
				addInventory(p, "Cyber Virus")
				fmt.Println("You bought Cyber Virus")
			}

		case 3:
			if len(p.Inventory) >= 10 {
				fmt.Println("Inventory is full")
			} else if p.Argent < 25 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 25
				addInventory(p, "Spellbook: Boule de Feu")
				fmt.Println("You bought Spellbook: Boule de Feu")
			}

		case 4:
			if len(p.Inventory) >= 10 {
				fmt.Println("Inventory is full")
			} else if p.Argent < 4 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 4
				addInventory(p, "Fourrure de Loup")
				fmt.Println("You bought Fourrure de Loup")
			}

		case 5:
			if len(p.Inventory) >= 10 {
				fmt.Println("Inventory is full")
			} else if p.Argent < 7 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 7
				addInventory(p, "Peau de Troll")
				fmt.Println("You bought Peau de Troll")
			}

		case 6:
			if len(p.Inventory) >= 10 {
				fmt.Println("Inventory is full")
			} else if p.Argent < 3 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 3
				addInventory(p, "Cuir de Sanglier")
				fmt.Println("You bought Cuir de Sanglier")
			}

		case 7:
			if len(p.Inventory) >= 10 {
				fmt.Println("Inventory is full")
			} else if p.Argent < 1 {
				fmt.Println("Not enough credits")
			} else {
				p.Argent -= 1
				addInventory(p, "Plume de Corbeau")
				fmt.Println("You bought Plume de Corbeau")
			}

		case 8:
			return

		default:
			fmt.Println("\nInvalid choice")
		}
	}
}

// MORT DU PERSONNAGE

func isDead(p *Character) bool {
	if p.HP <= 0 {
		fmt.Println("\n===================================")
		fmt.Println("          SYSTEM FAILURE")
		fmt.Println("              WASTED")
		fmt.Println("===================================")
		fmt.Println("You died in the streets of Neon City...")
		p.HP = p.MaxHP / 2
		fmt.Println("\nSystem reboot in progress...")
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

// MENU PRINCIPAL

func mainMenu(p *Character) {
	for {
		fmt.Println("\n===================================")
		fmt.Println("            NEON CITY")
		fmt.Println("===================================")
		fmt.Println("1. Display character profile")
		fmt.Println("2. Access inventory")
		fmt.Println("3. Cybernetic Workshop")
		fmt.Println("4. Quit")
		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			displayInfo(p)
		case 2:
			inventoryMenu(p)
		case 3:
			merchant(p)
		case 4:
			fmt.Println("\nExiting Neon City system...")
			return
		default:
			fmt.Println("\nInvalid choice, please try again.")
		}
	}
}

// MAIN
func main() {
	fmt.Println("===================================")
	fmt.Println("            NEON CITY")
	fmt.Println("===================================")
	fmt.Println("\nWelcome to Neon City, year 2087.")
	fmt.Println("Create your cybernetic character.\n")
	// On utilise maintenant la fonction de création
	character := characterCreation()
	fmt.Println("\n===================================")
	fmt.Println("       CHARACTER CREATED")
	fmt.Println("===================================")
	fmt.Println("Welcome,", character.Name)
	fmt.Println("Class:", character.Class)
	// Lancement du jeu
	mainMenu(&character)
}