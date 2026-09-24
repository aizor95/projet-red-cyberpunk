package main

import "fmt"

func main() {
	printHeader("NEON CITY // 2087")
	fmt.Println("La ville tourne encore. Les corpos aussi.")
	fmt.Println("Il est temps de creer ton personnage.")

	character := characterCreation()

	printPanel("Connexion etablie")
	fmt.Println("Bienvenue,", character.Name)
	fmt.Println("Classe :", character.Class)

	mainMenu(&character)
}
