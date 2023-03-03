package main

import "fmt"

type Car struct {
	carType   string
	carFuelIn float64
}

func (c Car) distance() float64 {

	return c.carFuelIn / 1.5

}

func main() {

	car1 := Car{"Toyota", 15}
	car2 := Car{"Honda", 10}
	car3 := Car{"Suzuki", 25}

	fmt.Println(car1.distance())
	fmt.Println(car2.distance())
	fmt.Println(car3.distance())

}
