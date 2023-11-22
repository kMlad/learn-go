package fundamentals

import "fmt"

func Variables() {
	const DaysInYear = 365

	var favoriteColor string = "black"
	fmt.Println("My favorite color is", favoriteColor)

	var birthYear, age = 1998, 25
	fmt.Println("I'm born in ", birthYear, " and I'm ", age, " years old.")

	var (
		firstInitial = "K"
		lastInitial  = "M"
	)
	fmt.Println("I'm ", firstInitial, ".", lastInitial, ".")

	var ageInDays int
	ageInDays = age * DaysInYear
	fmt.Println("I am ", ageInDays, " days old.")
}
