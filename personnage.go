package main

import (
	"fmt"
	"time"
	"unicode"
	"strings"
)

type Character struct {
	Name string
	Class string
	Level int
	HP int
	MaxHP int
	Energy int
	Attack int
	Defense int
	Argent int
	Equipment Equipment
	Inventory []string
	Skill     []string
}

type Equipment struct {
	head EquipmentItem
	torse EquipmentItem
	feet EquipmentItem
}

type EquipmentItem struct {
	Name string
	Slot string
	MaxHPbonus int
	Attack int
	Defense int
	Energy int
}

type MerchantItem struct {
	Name  string
	Price int
}

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
	case "Mercenary":
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

func accessInventory(p Character) []string {
	fmt.Println("\n=== STORAGE UNIT (INVENTORY) ===")

	if len(p.Inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return p.Inventory
	}

	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	return p.Inventory
}

func takePot(p *Character) {
	index := -1

	for i, item := range p.Inventory {
		if item == "Stimpack" {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Error: No Stimpack available in your inventory!")
		return
	}

	p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)

	p.HP += 50
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}

	fmt.Println(">>> Stimpack used! (+50 HP)")
	fmt.Printf("Current body integrity: %d / %d HP\n", p.HP, p.MaxHP)
}

func mainMenu(p *Character) {
	loop := true

	for loop {
		var choice int

		fmt.Println("\n===========================")
		fmt.Println("         MAIN MENU         ")
		fmt.Println("===========================")
		fmt.Println("1. Display character profile")
		fmt.Println("2. Access inventory")
		fmt.Println("3. Cybernetic Workshop")
		fmt.Println("4. Quit")
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\n--- CHARACTER PROFILE ---")
			fmt.Println("Name:", p.Name)
			fmt.Println("Class:", p.Class)
			fmt.Printf("Health: %d/%d HP\n", p.HP, p.MaxHP)
			fmt.Println("Attack:", p.Attack)
			fmt.Println("Skills:", p.Skill)

		case 2:
			accessInventory(*p)

			var invChoice int
			fmt.Println("\n[1] Use Stimpack | [2] Use Cyber Virus | [3] Read Spellbook | [4] Back")
			fmt.Print("Choice: ")
			fmt.Scanln(&invChoice)

			if invChoice == 1 {
				takePot(p)
			} else if invChoice == 2 {
				poisonPot(p)
			} else if invChoice == 3 {
				spellBook(p)
			}

		case 3:
			merchant(p)

		case 4:
			fmt.Println("Exiting Neon City system...")
			loop = false

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addInventory(p *Character, item string) {
	if isInventoryFull(p) {
		fmt.Println("\n[!] WARNING: Storage capacity reached! Limit is 10 items.")
		fmt.Printf("Could not acquire: %s\n", item)
		return
	}

	p.Inventory = append(p.Inventory, item)
	fmt.Printf(">> Added to storage: %s (%d/10 slots occupied)\n", item, len(p.Inventory))
}

func merchant(p *Character) {
	var choice int

	fmt.Println("\n==================================")
	fmt.Println("   CYBERNETIC WORKSHOP (MERCHANT) ")
	fmt.Println("==================================")
	fmt.Println("1. Stimpack (Free)")
	fmt.Println("2. Cyber Virus (Free)")
	fmt.Println("3. Spellbook: Boule de Feu (Free)")
	fmt.Println("4. Back to main menu")
	fmt.Print("\nYour choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		addInventory(p, "Stimpack")
	case 2:
		addInventory(p, "Cyber Virus")
	case 3:
		addInventory(p, "Spellbook: Boule de Feu")
	case 4:
		fmt.Println("Returning to main menu...")
	default:
		fmt.Println("Invalid choice.")
	}
}

func isDead(p *Character) bool {
	if p.HP <= 0 {
		fmt.Println("\n==================================")
		fmt.Println("         SYSTEM FAILURE           ")
		fmt.Println("             WASTED               ")
		fmt.Println("==================================")
		fmt.Println("You died in the streets of Neon City...")

		p.HP = p.MaxHP / 2

		fmt.Println("System reboot in progress...")
		fmt.Printf("You have been revived with %d / %d HP.\n", p.HP, p.MaxHP)
		fmt.Println("==================================")
		return true
	}
	return false
}

func poisonPot(p *Character) {
	fmt.Println("\n[!] Warning: Cyber Virus activated! Taking damage over time...")

	for tick := 1; tick <= 3; tick++ {
		time.Sleep(1 * time.Second)
		p.HP -= 10

		fmt.Printf("Tick %d/3 - HP: %d / %d\n", tick, p.HP, p.MaxHP)

		if isDead(p) {
			break
		}
	}
}

func spellBook(p *Character) {
	for _, s := range p.Skill {
		if s == "Boule de Feu" {
			fmt.Println("You have already learned the spell: Boule de Feu!")
			return
		}
	}

	p.Skill = append(p.Skill, "Boule de Feu")
	fmt.Println(">>> New spell learned: Boule de Feu!")
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
	lower := strings.ToLower(name)
	return strings.ToUpper(string(lower[0])) + lower[1:]
}

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
		fmt.Println("Invalid input! Name must contain letters only.")
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