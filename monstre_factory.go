package main

import "fmt"

func createMonster(name string) Monster {
	switch name {
	case "carcasse de robot":
		return initRobot()
	case "Jack-8":
		return initJack8()
	case "Drone de Surveillance":
		return initDrone()
	case "Robot Industriel":
		return initRobotIndustriel()
	default:
		fmt.Println("Erreur : monstre inconnu :", name)
		return Monster{}
	}
}
