package main

type Drop struct {
	ItemName string
	Chance   int
}

type Monster struct {
	Name string
	HP int
	MaxHP int
	Attack int
	Initiative int
	Poisoned bool
	PoisonTurns int
	XPReward int
	Drops []Drop
}

func initRobot() Monster {
	maxHP := 40
	return Monster{
		Name: "carcasse de robot",
		HP: maxHP,
		MaxHP: maxHP,
		Attack: 5,
		Initiative: 6,
		XPReward: 25,
		Drops: []Drop{
			{
				ItemName: "Cuir Synthétique",
				Chance: 40,
			},
		},
	}
}

func initJack8() Monster {
	return Monster{
		Name: "Jack-8",
		HP: 120,
		MaxHP: 120,
		Attack: 15,
		Initiative: 4,
		XPReward: 30,
		Drops: []Drop{
			{
				ItemName: "Peau de Cyborg",
				Chance: 70,
			},
			{
				ItemName: "Fibre Optique",
				Chance: 45,
			},
		},
	}
}

func initDrone() Monster {
	return Monster{
		Name: "Drone de Surveillance",
		HP: 50,
		MaxHP: 50,
		Attack: 8,
		Initiative: 15,
		XPReward: 15,
		Drops: []Drop{
			{
				ItemName: "Puce de Données",
				Chance: 60,
			},
		},
	}
}

func initRobotIndustriel() Monster {
	return Monster{
		Name:       "Robot Industriel",
		HP:         90,
		MaxHP:      90,
		Attack:     12,
		Initiative: 6,
		XPReward:   20,
		Drops: []Drop{
			{
				ItemName: "Fibre Optique",
				Chance: 55,
			},
		},
	}
}