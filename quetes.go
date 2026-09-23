package main

import "fmt"

type Quest struct {
	ID                 string
	Name               string
	Description        string
	ObjectiveType      string
	Target             string
	TargetLocation     string
	TargetNPC          string
	RequiredAmount     int
	RewardCredits      int
	RewardXP           int
	RewardBadge        int
	PrerequisiteQuests []string
	MinBadge           int
	Repeatable         bool
	StartItem          string
}

var quests = []Quest{
	// DÉBUT DU JEU
	{
		ID: "Q001",
		Name: "Nettoyage des ruelles",
		Description: "Éliminez 3 carcasses de robots dans la zone d'entraînement.",
		ObjectiveType: "kill",
		Target: "carcasse de robot",
		RequiredAmount: 3,
		RewardCredits: 3,
		RewardXP: 20,
		RewardBadge: 1,
		MinBadge: 0,
		Repeatable: false,
	},
	{
		ID: "Q002",
		Name: "Livraison express",
		Description: "Livrez un colis au contact Kira dans le Quartier Néon.",
		ObjectiveType: "delivery",
		Target: "Kira",
		TargetLocation: "Quartier Néon",
		TargetNPC: "Kira",
		RequiredAmount: 1,
		RewardCredits: 5,
		RewardXP: 25,
		MinBadge: 1,
		PrerequisiteQuests: []string{
			"Q001",
		},
		StartItem: "Colis : Quartier Néon",
	},
	{
		ID: "Q003",
		Name: "Chasse au cyborg",
		Description: "Éliminez 3 Jack-8 dans les zones dangereuses de Neon City.",
		ObjectiveType: "kill",
		Target: "Jack-8",
		RequiredAmount:3,
		RewardCredits: 8,
		RewardXP: 40,
		MinBadge: 1,
		PrerequisiteQuests: []string{
			"Q001",
		},
	},
	{
		ID: "Q004",
		Name: "Piratage industriel",
		Description: "Piratez le terminal industriel de la Zone industrielle.",
		ObjectiveType: "hack",
		Target: "Terminal industriel",
		RequiredAmount: 1,
		RewardCredits: 6,
		RewardXP: 35,
		MinBadge: 1,
		PrerequisiteQuests: []string{
			"Q001",
		},
	},
	// BADGE NIVEAU 2
	{
		ID: "Q005",
		Name: "Épreuve d'autorisation",
		Description: "Battez 1 Robot Industriel pour prouver votre capacité à pénétrer dans les secteurs corporatifs.",
		ObjectiveType: "kill",
		Target: "Robot Industriel",
		RequiredAmount: 1,
		RewardCredits: 10,
		RewardXP: 50,
		RewardBadge: 2,
		MinBadge: 1,
		PrerequisiteQuests: []string{
			"Q002",
			"Q003",
			"Q004",
		},
	},
	// APRÈS BADGE 2
	{
		ID: "Q006",
		Name: "Chasse aux drones",
		Description: "Éliminez 5 Drones de Surveillance.",
		ObjectiveType: "kill",
		Target: "Drone de Surveillance",
		RequiredAmount: 5,
		RewardCredits: 8,
		RewardXP: 50,
		MinBadge: 2,
		PrerequisiteQuests: []string{
			"Q005",
		},
	},
	{
		ID: "Q007",
		Name: "Livraison corporative",
		Description: "Livrez un colis à l'Agent Vega dans le Quartier corporatif.",
		ObjectiveType: "delivery",
		Target: "Agent Vega",
		TargetLocation: "Quartier corporatif",
		TargetNPC: "Agent Vega",
		RequiredAmount: 1,
		RewardCredits: 12,
		RewardXP: 60,
		MinBadge: 2,
		PrerequisiteQuests: []string{
			"Q005",
		},
		StartItem: "Colis : Quartier corporatif",
	},
	{
		ID: "Q008",
		Name: "Sécurité compromise",
		Description: "Piratez le terminal de sécurité corporatif.",
		ObjectiveType: "hack",
		Target: "Terminal de sécurité corporatif",
		RequiredAmount: 1,
		RewardCredits: 10,
		RewardXP: 50,
		MinBadge: 2,
		PrerequisiteQuests: []string{
			"Q005",
		},
	},
	// BADGE NIVEAU 3
	{
		ID: "Q009",
		Name: "Épreuve finale",
		Description: "Éliminez 2 Jack-8 pour obtenir l'autorisation d'accès aux zones interdites.",
		ObjectiveType: "kill",
		Target: "Jack-8",
		RequiredAmount: 2,
		RewardCredits: 15,
		RewardXP: 80,
		RewardBadge: 3,
		MinBadge: 2,
		PrerequisiteQuests: []string{
			"Q006",
			"Q007",
			"Q008",
		},
	},
	// QUÊTES RÉPÉTABLES
	{
		ID: "R001",
		Name: "Entraînement continu",
		Description: "Éliminez 5 carcasses de robots.",
		ObjectiveType: "kill",
		Target: "carcasse de robot",
		RequiredAmount: 5,
		RewardCredits: 3,
		RewardXP: 15,
		MinBadge: 0,
		Repeatable: true,
	},
	{
		ID: "R002",
		Name: "Chasse aux cyborgs",
		Description: "Éliminez 2 Jack-8.",
		ObjectiveType: "kill",
		Target: "Jack-8",
		RequiredAmount: 2,
		RewardCredits: 6,
		RewardXP: 25,
		MinBadge: 1,
		Repeatable: true,
	},
	{
		ID: "R003",
		Name: "Surveillance aérienne",
		Description: "Éliminez 3 Drones de Surveillance.",
		ObjectiveType: "kill",
		Target: "Drone de Surveillance",
		RequiredAmount: 3,
		RewardCredits:  8,
		RewardXP: 30,
		MinBadge: 2,
		Repeatable: true,
	},
	{
		ID: "R004",
		Name: "Nettoyage industriel",
		Description: "Éliminez 2 Robots Industriels.",
		ObjectiveType: "kill",
		Target: "Robot Industriel",
		RequiredAmount: 2,
		RewardCredits: 12,
		RewardXP: 40,
		MinBadge: 3,
		Repeatable: true,
	},
}

// RECHERCHE
func getQuest(id string) *Quest {
	for i := range quests {
		if quests[i].ID == id {
			return &quests[i]
		}
	}
	return nil
}

func getActiveQuest(p *Character) *Quest {
	if p.ActiveQuestID == "" {
		return nil
	}
	return getQuest(p.ActiveQuestID)
}

// QUÊTES TERMINÉES
func questCompleted(p *Character, questID string) bool {
	for _, id := range p.CompletedQuests {
		if id == questID {
			return true
		}
	}
	return false
}

func prerequisitesCompleted(p *Character, quest Quest) bool {
	for _, prerequisite := range quest.PrerequisiteQuests {
		if !questCompleted(p, prerequisite) {
			return false
		}
	}
	return true
}

// QUÊTES DISPONIBLES
func availableQuests(p *Character) []Quest {
	var available []Quest
	for _, quest := range quests {
		if quest.MinBadge > p.BadgeLevel {
			continue
		}
		if !prerequisitesCompleted(p, quest) {
			continue
		}
		if !quest.Repeatable && questCompleted(p, quest.ID) {
			continue
		}
		if quest.ID == p.ActiveQuestID {
			continue
		}
		available = append(available, quest)
	}
	return available
}

// AFFICHAGE
func showQuestDetails(quest Quest, progress int) {
	fmt.Println("\n===================================")
	fmt.Println("              QUÊTE")
	fmt.Println("===================================")
	fmt.Println("ID :", quest.ID)
	fmt.Println("Nom :", quest.Name)
	fmt.Println("\n", quest.Description)
	fmt.Printf(
		"\nProgression : %d/%d\n",
		progress,
		quest.RequiredAmount,
	)
	fmt.Println("\nRécompenses :")
	fmt.Println("+", quest.RewardCredits, "crédits")
	fmt.Println("+", quest.RewardXP, "XP")
	if quest.RewardBadge > 0 {
		fmt.Println("+ Badge niveau", quest.RewardBadge)
	}
	if quest.Repeatable {
		fmt.Println("Quête répétable : oui")
	}
}

// ACCEPTATION
func acceptQuest(p *Character, quest Quest) bool {
	if p.ActiveQuestID != "" {
		fmt.Println("\nVous avez déjà une quête active.")
		return false
	}
	if quest.StartItem != "" {
		if isInventoryFull(p) {
			fmt.Println("\nInventaire plein.")
			fmt.Println("Impossible de prendre le colis.")
			return false
		}
	}
	p.ActiveQuestID = quest.ID
	p.QuestProgress = 0
	if quest.StartItem != "" {
		addInventory(p, quest.StartItem)
	}
	fmt.Println("\n===================================")
	fmt.Println("          QUÊTE ACCEPTÉE")
	fmt.Println("===================================")
	fmt.Println(quest.Name)
	fmt.Println(quest.Description)
	return true
}

// ABANDON
func abandonQuest(p *Character) {
	quest := getActiveQuest(p)
	if quest == nil {
		return
	}
	if quest.StartItem != "" {
		removeInventoryItem(p, quest.StartItem)
	}
	fmt.Println("\nQuête abandonnée :", quest.Name)
	p.ActiveQuestID = ""
	p.QuestProgress = 0
}

// PROGRESSION
func updateQuestProgress(
	p *Character,
	objectiveType string,
	target string,
) {
	quest := getActiveQuest(p)
	if quest == nil {
		return
	}
	if quest.ObjectiveType != objectiveType {
		return
	}
	if quest.Target != target {
		return
	}
	p.QuestProgress++
	if p.QuestProgress > quest.RequiredAmount {
		p.QuestProgress = quest.RequiredAmount
	}
	fmt.Printf(
		"\nProgression de quête : %s %d/%d\n",
		quest.Name,
		p.QuestProgress,
		quest.RequiredAmount,
	)
	if p.QuestProgress >= quest.RequiredAmount {
		completeActiveQuest(p)
	}
}

// OUTIL INVENTAIRE
func hasInventoryItem(p *Character, itemName string) bool {

	for _, item := range p.Inventory {

		if item == itemName {
			return true
		}
	}

	return false
}

// PIRATAGE
func hackTerminal(p *Character, terminalName string) {
	fmt.Println("\n===================================")
	fmt.Println("          PIRATAGE TERMINAL")
	fmt.Println("===================================")
	fmt.Println("Terminal ciblé :", terminalName)
	fmt.Println("Connexion au système...")
	updateQuestProgress(p, "hack", terminalName)
}

// RÉCOMPENSE
func completeActiveQuest(p *Character) {
	quest := getActiveQuest(p)
	if quest == nil {
		return
	}
	fmt.Println("\n===================================")
	fmt.Println("          QUÊTE TERMINÉE")
	fmt.Println("===================================")
	fmt.Println(quest.Name)
	if quest.StartItem != "" {
		removeInventoryItem(p, quest.StartItem)
	}
	if quest.RewardCredits > 0 {
		p.Argent += quest.RewardCredits
		fmt.Println(
			"+",
			quest.RewardCredits,
			"crédits",
		)
	}
	if quest.RewardXP > 0 {
		gainExperience(p, quest.RewardXP)
		fmt.Println(
			"+",
			quest.RewardXP,
			"XP de quête",
		)
	}
	if quest.RewardBadge > p.BadgeLevel {
		p.BadgeLevel = quest.RewardBadge
		fmt.Println(
			"\n>>> BADGE NIVEAU",
			quest.RewardBadge,
			"OBTENU !",
		)
		switch quest.RewardBadge {
		case 1:
			fmt.Println("La Zone industrielle est maintenant accessible.")
		case 2:
			fmt.Println("Le Quartier corporatif est maintenant accessible.")
		case 3:
			fmt.Println("La Zone interdite est maintenant accessible.")
		}
	}
	if !quest.Repeatable {
		if !questCompleted(p, quest.ID) {
			p.CompletedQuests = append(
				p.CompletedQuests,
				quest.ID,
			)
		}
	}
	p.ActiveQuestID = ""
	p.QuestProgress = 0
	fmt.Println("\nSolde :", p.Argent, "crédits")
}