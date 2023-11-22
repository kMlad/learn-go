package fundamentals

import "fmt"

func calculateStageOfLife(age int) string {
	switch a := age; {
	case a == 0:
		return "Newborn"
	case a <= 3:
		return "Toddler"
	case a < 13:
		return "Child"
	case a < 18:
		return "Teenage"
	default:
		return "Adult"
	}
}

func Switch() {
	fmt.Println(calculateStageOfLife(0))
	fmt.Println(calculateStageOfLife(14))
	fmt.Println(calculateStageOfLife(10))
	fmt.Println(calculateStageOfLife(2))
	fmt.Println(calculateStageOfLife(20))
	fmt.Println(calculateStageOfLife(4))
}
