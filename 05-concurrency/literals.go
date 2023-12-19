package concurrency

import (
	"fmt"
	"unicode"
)

func Literals() {

	letters, digits, spaces := 0, 0, 0

	classifyRunes := func(str string) {
		for _, char := range str {
			if unicode.IsDigit(char) {
				digits++
			}
			if unicode.IsSpace(char) {
				spaces++
			}
			if unicode.IsLetter(char) {
				letters++
			}
		}
	}

	lines := []string{
		"There are",
		"52 letters,",
		"four digits and ",
		"11 spaces ",
		"in these lines of text!",
	}

	for _, line := range lines {
		classifyRunes(line)
	}

	fmt.Printf("There are %v letters, %v digits and %v spaces in this line of text!",
		letters,
		digits,
		spaces)
}
