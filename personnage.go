package main

import (
	"fmt"
	"time"
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
		Attack: attack,
		Defense: defense,
		Argent: 100,
		Equipment: Equipment{},
		Inventory: []string{},
	}
	return character
}

func displayInfo(player *Character) {
	fmt.Println("\n=== NEON CITY - PROFILE ===")
	fmt.Println("Nom:", player.Name)
	fmt.Println("Classe:", player.Class)
	fmt.Println("Niveau:", player.Level)
	fmt.Printf("Intégrité (HP): %d/%d\n", player.HP, player.MaxHP)
	fmt.Println("Énergie:", player.Energy)
	fmt.Println("Attaque:", player.Attack)
	fmt.Println("Défense:", player.Defense)
	fmt.Println("Crédits:", player.Argent)
	fmt.Println("Équipement:", player.Equipment)
	fmt.Println("Inventaire:", player.Inventory)
	fmt.Println("Compétences:", player.Skill)
	fmt.Println("===========================\n")
}

func takePot(player *Character) {
	if removeInventory(player, "Stimpack HP") {
		player.HP += 50
		if player.HP > player.MaxHP {
			player.HP = player.MaxHP
		}
		fmt.Printf("Stimpack utilisé ! HP actuels: %d/%d\n", player.HP, player.MaxHP)
	} else {
		fmt.Println("Aucun Stimpack HP dans l'inventaire !")
	}
}

func removeInventory(player *Character, item string) bool {
	for i := 0; i < len(player.Inventory); i++ {
		if player.Inventory[i] == item {
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func accessInventoryMenu(player *Character) {
	inInventory := true
	var choice int

	for inInventory {
		fmt.Println("\n=== NEON CITY - INVENTAIRE ===")
		fmt.Println("1. Voir les objets de l'inventaire")
		fmt.Println("2. Utiliser un Stimpack HP")
		fmt.Println("0. Retour au menu principal")
		fmt.Println("==============================")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nObjets actuels :", player.Inventory)
		case 2:
			takePot(player)
		case 0:
			
		default:
			fmt.Println("\n[ERREUR] Choix invalide.")
		}
	}
}

func mainMenu(player *Character) {
	running := true
	var choice int

	for running {
		fmt.Println("\n=== NEON CITY - TERMINAL ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Quitter")
		fmt.Println("=============================")
		fmt.Print("Faites votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displayInfo(player)
		case 2:
			accessInventoryMenu(player)
		case 3:
			fmt.Println("\n[SYSTEME] Fermeture de la connexion... Au revoir, mercenaire.")
			running = false 
		default:
			fmt.Println("\n[ERREUR] Commande inconnue dans le réseau.")
		}
	}
}
func merchantMenu(player *Character) {
	items := []MerchantItem{
		{Name: "Stimpack HP", Price: 20},
		{Name: "Virus DoT", Price: 30},
		{Name: "Programme: Boule de Feu", Price: 50},
		{Name: "Potion de poison", Price: 15},
	}

	fmt.Println("\n--- MARCHAND NEON CITY ---")
	for i, item := range items {
		fmt.Printf("%d. %s - %d Crédits\n", i+1, item.Name, item.Price)
	}
	fmt.Println("4. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1, 2, 3:
		item := items[choice-1]
		
		if player.Argent < item.Price {
			fmt.Println("Crédits insuffisants !")
			return
		}

		player.Argent -= item.Price
		player.Inventory = append(player.Inventory, item.Name)
		fmt.Printf("Achat réussi : %s\n", item.Name)

	case 4:
		return

	default:
		fmt.Println("Choix invalide.")
	}
}

func isDead(player *Character) bool {
	if player.HP <= 0 {
		fmt.Println("\n[WASTED] Vous êtes mort...")
		player.HP = player.MaxHP / 2
		fmt.Printf("Ressuscité ! Vos PV ont été restaurés à 50%% (%d/%d HP).\n", player.HP, player.MaxHP)
		return true
	}
	return false
}
func poisonPot(player *Character) {
	if !removeInventory(player, "Potion de poison") {
		fmt.Println("No Poison Potion in your inventory!")
		return
	}

	fmt.Println("\n[POISON] You drank a Poison Potion! Taking damage...")

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		player.HP -= 10
		if player.HP < 0 {
			player.HP = 0
		}

		fmt.Printf("Second %d/3: Current HP: %d/%d\n", i, player.HP, player.MaxHP)
		if isDead(player) {
			break
		}
	}
}

func spellBook(player *Character) {
	spellName := "Boule de Feu"

	for _, s := range player.Skill {
		if s == spellName {
			fmt.Printf("Vous connaissez déjà le sort %s !\n", spellName)
			return
		}
	}

	if removeInventory(player, "Livre de Sort : Boule de Feu") {
		player.Skill = append(player.Skill, spellName)
		fmt.Printf("Succès ! Vous avez appris le sort : %s\n", spellName)
	} else {
		fmt.Println("Vous n'avez pas de Livre de Sort dans votre inventaire !")
	}
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