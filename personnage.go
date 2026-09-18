package main

import "fmt"

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

func initCharacter(name string,class string) Character {
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
	character := Character{
		Name: name,
		Class: class,
		Level: 1,
		HP: maxHP,
		MaxHP: maxHP,
		Energy: energy,
		Defense: defense,
		Argent: 100,
		Equipment: Equipment{},
		Inventory: []string{},
	}
	return character
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
}
