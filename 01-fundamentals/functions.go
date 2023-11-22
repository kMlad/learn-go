package fundamentals

import (
	"fmt"
	"math/rand"
)

func greetPerson(name string) {
	fmt.Println("Hello ", name, "!")
}

func addNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func rng() int {
	return rand.Intn(100)
}

func twoRng() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}

func Functions() {
	greetPerson("Kiril")
	fmt.Println(addNumbers(1, 2, 3))
	fmt.Println(rng())
	fmt.Println(twoRng())
}
