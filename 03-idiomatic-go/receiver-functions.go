package idiomaticgo

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) changeHealth(newHealth uint) {
	if newHealth > player.maxHealth {
		// panic(fmt.Sprintf("%v exceeds the player's max health.", newHealth))
		player.health = player.maxHealth
	} else {
		player.health = newHealth
	}
}

func (player *Player) changeEnergy(newEnergy uint) {
	if newEnergy > player.maxEnergy {
		// panic(fmt.Sprintf("%v exceeds the player's max Energy.", newEnergy))
		player.energy = player.maxEnergy
	} else {
		player.energy = newEnergy
	}
}

func (player *Player) printPlayerStatus() {
	fmt.Println("----------------------------------------")
	fmt.Println("Current player health: ", player.health)
	fmt.Println("Current player energy: ", player.energy)
	fmt.Println("----------------------------------------")
}

func ReceiverFunctions() {
	player := Player{
		name:      "Max",
		health:    20,
		maxHealth: 100,
		energy:    20,
		maxEnergy: 100,
	}
	player.printPlayerStatus()

	player.changeHealth(120)
	player.printPlayerStatus()

	player.changeEnergy(80)
	player.printPlayerStatus()

}
