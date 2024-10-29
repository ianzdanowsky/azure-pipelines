package main

import "fmt"

type car struct {
	name string
	year string
}

func (c *car) getCarYear() {
	fmt.Println(c.year)
}

func main() {
	newCar := car{name: "Toyota", year: "1994"}
	fmt.Println("Hello world")
	newCar.getCarYear()
}
