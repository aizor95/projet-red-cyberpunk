package main

type Monster struct {
	Name   string
	HP     int
	MaxHP  int
	Attack int
}

func initGoblin() Monster {
	maxHP := 40
	return Monster{
		Name:   "carcasse de robot",
		MaxHP:  maxHP,
		HP:     maxHP,
		Attack: 5,
	}
}
