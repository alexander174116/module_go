// ...это возможность работать с объектами разных типов через один и тот же интерфейс
// определим интерфейс
package main

import "fmt"

type Greeter interface {
	Greet()
}

// создадим три разных типа, реализующих интерфейс
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Привет, меня зовут", p.Name)
}

type Dog struct {
	Name string
}

func (d Dog) Greet() {
	fmt.Println("Гав-гав! Я", d.Name)
}

type Robot struct {
	Name string
}

func (r Robot) Greet() {
	fmt.Println("Здравствуйте. Модель", r.Name, "включена")
}

// напишем универсальную функцию, которая принимает интерфейс и вызывает у него метод:
func SayHello(g Greeter) {
	g.Greet()
}

// проверка типа во время выполнения:
// конструкции type assertion или type swith позволяют узнать, какой конкретно тип стоит за интерфейсной переменной:
func Describe(g Greeter) {
	switch v := g.(type) {
	case Person:
		fmt.Println("Это человек по имени", v.Name)
	case Dog:
		fmt.Println("Это собака по кличке", v.Name)
	case Robot:
		fmt.Println("Это робот модели", v.Name)
	default:
		fmt.Println("Неизвестный тип")
	}
}

// теперь мы можем сделать следующее в main:
/*
func main() {
	p := Person{Name: "Anna"}
	d := Dog{Name: "Sharik"}
	r := Robot{Name: "XJ-9"}
	SayHello(p)
	SayHello(d)
	SayHello(r)
}
*/

// go позволяет создать срез, содержащий объекты разных типов
// если они все реализуют один и тот же интерфейс
func main() {
	p := Person{Name: "Anna"}
	d := Dog{Name: "Sharik"}
	r := Robot{Name: "XJ-9"}
	greeters := []Greeter{p, d, r}
	for _, g := range greeters {
		Describe(g)
		g.Greet()
	}
	// SayHello(p)
	// SayHello(d)
	// SayHello(r)
}
