package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	res := r.Width * r.Height
	// fmt.Println("Площадь прямоугольника:", res)
	return res
}

func (c Circle) Area() float64 {
	res := math.Pi * math.Pow(c.Radius, 2)
	// fmt.Println("Площадь круга:", res)
	return res
}
func TotalArea(s []Shape) float64 {
	res := 0.0
	for _, g := range s {
		res += g.Area()
	}
	return res
}

func main() {
	rect_a := Rectangle{Width: 5.15, Height: 10.11}
	rect_b := Rectangle{Width: 5.15, Height: 10.11}
	rect_c := Rectangle{Width: 5.15, Height: 10.11}
	circ_a := Circle{Radius: 3.33}
	circ_b := Circle{Radius: 3.33}
	circ_c := Circle{Radius: 3.33}
	srez := []Shape{rect_a, rect_b, rect_c, circ_a, circ_b, circ_c}
	fmt.Println("Площадь всех фигур:", TotalArea(srez))
}
