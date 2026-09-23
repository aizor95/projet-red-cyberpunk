package main

import "fmt"

type Location struct {
	Name string
	Description string
	Connections []string
	Merchant bool
	Forgeron bool
	Training bool
	Enemies []string
	RequiredBadge int
}

func initMap() map[string]Location {
	return map[string]Location{
		"Place centrale": {
			Name: "Place centrale",
			Description: "Le cœur de Neon City. Une immense place entourée de néons.",
			Connections: []string{
				"Marché noir",
				"Ruelles",
				"Station de métro",
			},
		},
		"Marché noir": {
			Name: "Marché noir",
			Description: "Un marché clandestin où circulent marchandises et technologies.",
			Connections: []string{
				"Place centrale",
			},
			Merchant: true,
			Forgeron: true,
		},
		"Ruelles": {
			Name: "Ruelles",
			Description: "Des ruelles dangereuses où les affrontements sont fréquents.",
			Connections: []string{
				"Place centrale",
			},
			Training: true,
		},
		"Station de métro": {
			Name: "Station de métro",
			Description: "Une station permettant de rejoindre différents quartiers de Neon City.",
			Connections: []string{
				"Place centrale",
				"Zone industrielle",
				"Quartier Néon",
				"Quartier corporatif",
			},
		},
		"Zone industrielle": {
			Name: "Zone industrielle",
			Description: "Une zone remplie d'usines abandonnées, de machines et de déchets électroniques.",
			Connections: []string{
				"Station de métro",
				"Les Bas-Fonds",
			},
			Enemies: []string{
				"carcasse de robot",
				"Drone de Surveillance",
				"Jack-8",
				"Robot Industriel",
			},
			RequiredBadge: 1,
		},
		"Les Bas-Fonds": {
			Name: "Les Bas-Fonds",
			Description: "Un quartier dangereux contrôlé par différents gangs.",
			Connections: []string{
				"Zone industrielle",
				"Réseau souterrain",
			},
		},
		"Quartier Néon": {
			Name: "Quartier Néon",
			Description: "Un quartier rempli de casinos, clubs et commerces luxueux.",
			Connections: []string{
				"Station de métro",
				"Quartier corporatif",
			},
		},
		"Quartier corporatif": {
			Name: "Quartier corporatif",
			Description: "Le quartier des grandes corporations et des bâtiments ultra-sécurisés.",
			Connections: []string{
				"Quartier Néon",
				"Station de métro",
			},
			RequiredBadge: 2,
		},
		"Réseau souterrain": {
			Name: "Réseau souterrain",
			Description: "Un ancien réseau souterrain reliant les zones les plus dangereuses de la ville.",
			Connections: []string{
				"Les Bas-Fonds",
				"Zone interdite",
			},
		},
		"Zone interdite": {
			Name: "Zone interdite",
			Description: "Une zone ultra-sécurisée où se cachent les installations les plus dangereuses de Neon City.",
			Connections: []string{
				"Réseau souterrain",
			},
			RequiredBadge: 3,
		},
	}
}

func explore(p *Character) {
	worldMap := initMap()
	for {
		location := worldMap[p.CurrentLocation]
		fmt.Println("\n===================================")
		fmt.Println("             NEON CITY")
		fmt.Println("===================================")
		fmt.Println("Position :", location.Name)
		fmt.Println(location.Description)
		fmt.Println("\n---------------- ACTIONS ----------------")
		var actions []string
		if location.Merchant {
			actions = append(actions, "Marchand")
		}
		if location.Forgeron {
			actions = append(actions, "Forgeron")
		}
		if location.Training {
			actions = append(actions, "Zone d'entraînement")
		}
		if len(location.Enemies) > 0 {
			actions = append(actions, "Explorer la zone")
		}
		for i, action := range actions {
			fmt.Printf("%d. %s\n", i+1, action)
		}
		fmt.Println("\n---------------- DÉPLACEMENT ----------------")
		for i, connection := range location.Connections {
			destination := worldMap[connection]
			if p.BadgeLevel >= destination.RequiredBadge {
				fmt.Printf("%d. Aller vers %s\n",i+1+len(actions),
			connection,)
			} else {
				fmt.Printf("%d. %s [BLOQUÉ - Badge %d requis]\n",i+1+len(actions),connection,destination.RequiredBadge,)
			}
		}
		fmt.Println("0. Quitter l'exploration")
		var choice int
		fmt.Print("\nVotre choix : ")
		fmt.Scanln(&choice)
		if choice == 0 {
			return
		}
		if choice >= 1 && choice <= len(actions) {
			action := actions[choice-1]
			switch action {
			case "Marchand":
				merchant(p)
			case "Forgeron":
				forgeron(p)
			case "Zone d'entraînement":
				trainingFight(p)
			case "Explorer la zone":
				randomEncounter(p, location)
			}
			continue
		}
		choice -= len(actions)
		if choice >= 1 && choice <= len(location.Connections) {
			newLocation := location.Connections[choice-1]
			destination := worldMap[newLocation]
			if p.BadgeLevel < destination.RequiredBadge {
				fmt.Println("\n===================================")
			fmt.Println("           ACCÈS REFUSÉ")
			fmt.Println("===================================")
			fmt.Println("Zone :", destination.Name)
			fmt.Printf("Badge niveau %d requis.\n",destination.RequiredBadge,)
			fmt.Printf("Votre badge actuel : niveau %d\n",p.BadgeLevel,)
			continue
		}
		p.CurrentLocation = newLocation
		fmt.Println("\nVous vous déplacez vers :", newLocation)
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}