package main

import (
	"fmt"
)

type Flyer interface {
	Fly()
}

type Bird struct {
	Name string
}

type Plane struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println("Птица по имени", b.Name, "летит!")
}

func (p Plane) Fly() {
	fmt.Println("Самолет модели", p.Name, "летит!")
}

func TestFlight(name Flyer) {
	name.Fly()
}

func main() {
	var f Flyer
	f = Bird{"Рио"}
	TestFlight(f)
	f = Plane{"БОЕИНГ 777"}
	TestFlight(f)
}
