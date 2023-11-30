package idiomaticgo

import "fmt"

type Calculator byte

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

func (calc Calculator) calculate(a, b int) int {
	switch calc {
	case Add:
		return a + b
	case Subtract:
		return a - b
	case Multiply:
		return a * b
	case Divide:
		return a / b
	}
	panic("Unhandled operation")
}

func Iota() {
	add := Calculator(Add)
	subtract := Calculator(Subtract)
	multiply := Calculator(Multiply)
	divide := Calculator(Divide)

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(subtract.calculate(10, 3)) // = 7

	fmt.Println(multiply.calculate(3, 3)) // = 9

	fmt.Println(divide.calculate(100, 2)) // = 50
}
