package types

import "fmt"

type Rectangle struct {
	length, width float32
}

func calculatePerimeter(rectangle Rectangle) float32 {
	return 2*rectangle.length + 2*rectangle.width
}

func calculateArea(rectangle Rectangle) float32 {
	return rectangle.length * rectangle.width
}

func printResults(rectangle Rectangle) {
	fmt.Println("The perimeter is: ", calculatePerimeter(rectangle))
	fmt.Println("The area is: ", calculateArea(rectangle))
}

func Structs() {
	rectangle1 := Rectangle{
		length: 23.2,
		width:  11.1,
	}
	rectangle2 := Rectangle{7, 11}
	rectangle3 := Rectangle{9, 11}
	rectangle4 := Rectangle{4.44, 44.4}
	rectangle5 := Rectangle{6.9, 42.0}

	printResults(rectangle1)
	printResults(rectangle2)
	printResults(rectangle3)
	printResults(rectangle4)
	printResults(rectangle5)
}
