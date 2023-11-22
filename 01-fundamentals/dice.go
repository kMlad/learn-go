package fundamentals

import (
	"fmt"
	"math/rand"
)

func roll(times, dice, sides int) string {
	total := 0

	for i := 0; i < times; i++ {
		for j, roll := 0, 0; j < dice; j++ {
			roll = rand.Intn(sides) + 1
			total += roll
		}
	}

	fmt.Println(total)
	if dice == 2 && total == 2 {
		return "Snake Eyes"
	}
	if total == 7 {
		return "Lucky 7"
	}
	if total%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func Dice() {
	fmt.Println(roll(1, 2, 2))
	fmt.Println(roll(1, 2, 6))
	fmt.Println(roll(2, 2, 3))
	fmt.Println(roll(4, 1, 3))
	fmt.Println(roll(5, 2, 6))
	fmt.Println(roll(7, 3, 7))
	fmt.Println(roll(2, 4, 8))
	fmt.Println(roll(8, 6, 99))
}
