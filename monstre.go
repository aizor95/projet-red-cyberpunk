package main

type Monster struct {
	Name string
	HP int
	MaxHP int
	Attack int
	Poisoned bool
	PoisonTurns int
}

func initGoblin() Monster {
	maxHP := 40
	return Monster{
		Name:"carcasse de robot",
		MaxHP:maxHP,
		HP:maxHP,
		Attack: 5,
	}
}
