package main

import "fmt"

// создаем интерфейс, задающий поведение
type Transporter interface {
	Move(passenger string)
}

// определим 2 типа
type Car struct {
	Brand string
}

func (c Car) Move(passenger string) {
	fmt.Printf("%s едет в машине %s\n", passenger, c.Brand)
}

type Bicycle struct {
	Color string
}

func (b Bicycle) Move(passenger string) {
	fmt.Printf("%s едет на велосипеде цвета %s\n", passenger, b.Color)
}

// напишем функцию, использующую интерфейс:
func StartRide(t Transporter, name string) {
	t.Move(name)
}

// main
func main() {
	car := Car{Brand: "Toyota"}
	bike := Bicycle{Color: "red"}
	StartRide(car, "OLEG")
	StartRide(bike, "IRINA")
}
