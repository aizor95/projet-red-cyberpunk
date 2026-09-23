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
	Terminal string
	NPCs []NPC
}

func initMap() map[string]Location {
	return map[string]Location{
		// PLACE CENTRALE
		"Place centrale": {
			Name: "Place centrale",
			Description:
				"Le cœur de Neon City. Une immense place entourée de néons.",
			Connections: []string{
				"Marché noir",
				"Ruelles",
				"Station de métro",
			},
			RequiredBadge: 0,
		},
		// MARCHÉ NOIR
		"Marché noir": {
			Name: "Marché noir",
			Description:
				"Un marché clandestin où circulent marchandises et technologies.",
			Connections: []string{
				"Place centrale",
			},
			Merchant: true,
			Forgeron: true,
			RequiredBadge: 0,
		},
		// RUELLES
		"Ruelles": {
			Name: "Ruelles",
			Description:
				"Des ruelles dangereuses où les affrontements sont fréquents.",
			Connections: []string{
				"Place centrale",
			},
			Training: true,
			RequiredBadge: 0,
		},
		// STATION DE MÉTRO
		"Station de métro": {
			Name: "Station de métro",
			Description:
				"Une station permettant de rejoindre différents quartiers de Neon City.",
			Connections: []string{
				"Place centrale",
				"Zone industrielle",
				"Les Bas-Fonds",
				"Quartier Néon",
				"Quartier corporatif",
			},
			RequiredBadge: 0,
		},
		// ZONE INDUSTRIELLE
		"Zone industrielle": {
			Name: "Zone industrielle",
			Description:
				"Une zone remplie d'usines abandonnées, de machines et de déchets électroniques.",
			Connections: []string{
				"Station de métro",
				"Les Bas-Fonds",
			},
			Enemies: []string{
				"Drone de Surveillance",
				"Jack-8",
				"Robot Industriel",
			},
			Terminal: "Terminal industriel",
			RequiredBadge: 1,
		},
		// BAS-FONDS
		"Les Bas-Fonds": {
			Name: "Les Bas-Fonds",
			Description:
				"Un quartier dangereux contrôlé par différents gangs.",
			Connections: []string{
				"Zone industrielle",
				"Réseau souterrain",
				"Station de métro",
			},
			RequiredBadge: 0,
		},
		// QUARTIER NÉON
		"Quartier Néon": {
			Name: "Quartier Néon",
			Description:
				"Un quartier rempli de casinos, clubs et commerces luxueux.",
			Connections: []string{
				"Station de métro",
				"Quartier corporatif",
			},
			RequiredBadge: 0,
			NPCs: []NPC{
				{
					Name: "Kira",
					Description: "Une informatrice spécialisée dans les échanges clandestins.",
					Dialogue: "Tu as le colis que je t'ai demandé ?",
				},
			},
		},
		// QUARTIER CORPORATIF
		"Quartier corporatif": {
			Name: "Quartier corporatif",
			Description:
				"Le quartier des grandes corporations et des bâtiments ultra-sécurisés.",
			Connections: []string{
				"Quartier Néon",
				"Station de métro",
			},
			Terminal: "Terminal de sécurité corporatif",
			RequiredBadge: 2,
			NPCs: []NPC{
				{
					Name: "Agent Vega",
					Description: "Un agent corporatif chargé de surveiller les échanges sensibles.",
					Dialogue: "Vous avez quelque chose pour moi ?",
				},
			},
		},
		// RÉSEAU SOUTERRAIN
		"Réseau souterrain": {
			Name: "Réseau souterrain",
			Description:
				"Un ancien réseau souterrain reliant les zones les plus dangereuses de la ville.",
			Connections: []string{
				"Les Bas-Fonds",
				"Zone interdite",
			},
			RequiredBadge: 0,
		},
		// ZONE INTERDITE
		"Zone interdite": {
			Name: "Zone interdite",
			Description:
				"Une zone ultra-sécurisée où se cachent les installations les plus dangereuses de Neon City.",
			Connections: []string{
				"Réseau souterrain",
			},
			Terminal: "Terminal de la zone interdite",
			RequiredBadge: 3,
		},
	}
}

// ACCÈS AUX ZONES
func canAccessLocation(
	p *Character,
	location Location,
) bool {
	return p.BadgeLevel >= location.RequiredBadge
}

// EXPLORATION
func explore(p *Character) {
	worldMap := initMap()
	for {
		location, exists := worldMap[p.CurrentLocation]
		if !exists {
			fmt.Println(
				"Erreur : position inconnue.",
			)
			p.CurrentLocation = "Place centrale"
			continue
		}
		fmt.Println("\n===================================")
		fmt.Println("             NEON CITY")
		fmt.Println("===================================")
		fmt.Println(
			"Position :",
			location.Name,
		)
		fmt.Println(
			location.Description,
		)
		// ACTIONS
		fmt.Println(
			"\n=================================== ACTIONS ===================================",
		)
		var actions []string
		if location.Merchant {
			actions = append(
				actions,
				"Marchand",
			)
		}
		if location.Forgeron {
			actions = append(
				actions,
				"Forgeron",
			)
		}
		if location.Training {
			actions = append(
				actions,
				"Zone d'entraînement",
			)
		}
		if len(location.Enemies) > 0 {
			actions = append(
				actions,
				"Explorer la zone",
			)
		}
		if location.Terminal != "" {
			actions = append(
				actions,
				"Pirater le terminal",
			)
		}
		for _, npc := range location.NPCs {
			actions = append(actions, "Parler à "+npc.Name)
		}
		for i, action := range actions {
			fmt.Printf(
				"%d. %s\n",
				i+1,
				action,
			)
		}
		// DÉPLACEMENT
		fmt.Println(
			"\n=================================== DÉPLACEMENT ===================================",
		)
		for i, connection := range location.Connections {
			destination := worldMap[connection]
			if canAccessLocation(
				p,
				destination,
			) {
				fmt.Printf(
					"%d. Aller vers %s\n",
					i+1+len(actions),
					connection,
				)
			} else {
				fmt.Printf(
					"%d. %s [BLOQUÉ - Badge %d requis]\n",
					i+1+len(actions),
					connection,
					destination.RequiredBadge,
				)
			}
		}
		// SORTIE
		fmt.Println(
			"0. Quitter l'exploration",
		)
		var choice int
		fmt.Print("\nVotre choix : ")
		fmt.Scanln(&choice)
		if choice == 0 {
			return
		}
		// ACTION
		if choice >= 1 &&
			choice <= len(actions) {
			action := actions[choice-1]
			switch action {
			case "Marchand":
				merchant(p)
			case "Forgeron":
				forgeron(p)
			case "Zone d'entraînement":
				trainingFight(p)
			case "Explorer la zone":
				randomEncounter(
					p,
					location,
				)
			case "Pirater le terminal":
				hackTerminal(
					p,
					location.Terminal,
				)
			case "Parler à Kira","Parler à Agent Vega" :
				for _, npc := range location.NPCs {
					if npc.Name == "Kira" {
						interactWithNPC(p, npc, location.Name)
					}
				}
			}
			continue
		}
		// DÉPLACEMENT
		moveChoice := choice - len(actions)
		if moveChoice < 1 ||
			moveChoice > len(location.Connections) {
			fmt.Println(
				"Choix invalide.",
			)
			continue
		}
		newLocationName :=
			location.Connections[moveChoice-1]
		destination :=
			worldMap[newLocationName]
		if !canAccessLocation(
			p,
			destination,
		) {
			fmt.Println(
				"\n===================================",
			)
			fmt.Println(
				"            ACCÈS REFUSÉ",
			)
			fmt.Println(
				"===================================",
			)
			fmt.Println(
				"Zone :",
				destination.Name,
			)
			fmt.Printf(
				"Badge niveau %d requis.\n",
				destination.RequiredBadge,
			)
			fmt.Printf(
				"Votre badge actuel : niveau %d\n",
				p.BadgeLevel,
			)
			continue
		}
		p.CurrentLocation =
			newLocationName
		fmt.Println(
			"\nVous vous déplacez vers :",
			newLocationName,
		)
		checkQuestAtLocation(
			p,
			newLocationName,
		)
	}
}