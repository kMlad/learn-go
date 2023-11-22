package fundamentals

import "fmt"

func isDivBy3(num int) bool {
	return num%3 == 0
}

func isDivBy5(num int) bool {
	return num%5 == 0
}

func playFizzBuzz(ceiling int) {
	for i := 1; i <= ceiling; i++ {
		switch num := i; {
		case isDivBy3(num) && isDivBy5(num):
			fmt.Println(num, " -> FizzBuzz")
		case isDivBy5(num):
			fmt.Println(num, " -> Buzz")
		case isDivBy3(num):
			fmt.Println(num, " -> Fizz")
		}
	}
}

func Loops() {
	playFizzBuzz(999)
}
