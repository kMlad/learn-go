package interfaces

import "fmt"

type Lifter interface {
	Lift()
}

type Motorcycle struct {
	modelName string
}
type Car struct {
	modelName string
}
type Truck struct {
	modelName string
}

func (bike Motorcycle) Lift() {
	fmt.Println("Taking ", bike.modelName, " to small lift.")
}
func (car Car) Lift() {
	fmt.Println("Taking ", car.modelName, " to standard lift.")
}
func (truck Truck) Lift() {
	fmt.Println("Taking ", truck.modelName, " to large lift.")
}

func liftVehicles(vehicles []Lifter) {
	for _, vehicle := range vehicles {
		vehicle.Lift()
	}
}

func Interfaces() {
	m2 := Car{"BMW M2"}
	mt07 := Motorcycle{"Yamaha MT-07"}
	cybertruck := Truck{"Tesla Cybertruck"}

	vehicles := []Lifter{m2, mt07, cybertruck}

	liftVehicles(vehicles)
}
