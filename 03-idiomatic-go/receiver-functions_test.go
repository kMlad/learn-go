package idiomaticgo

import "testing"

func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
	}
}

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.changeHealth(900)
	if player.health > player.maxHealth {
		t.Error("Player health exceeds max health!")
	}

}

func TestEnergy(t *testing.T) {
	player := newPlayer()
	player.changeEnergy(700)
	if player.energy > player.maxEnergy {
		t.Error("Player health exceeds max energy!")
	}
}
