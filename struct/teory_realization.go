package main

import "fmt"

type Speaker interface { //объявляем интрефейс
	Speak()
	// MyAge()
}

type Human struct { //объявляем структуру
	Name string
	// Age  int
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println("Гав! Я", d.Name)
}

func (h Human) Speak() { //реализуем метод Speak
	fmt.Println("Привет, меня зовут", h.Name)
}

// func (h Human) MyAge() {
// 	fmt.Println("Мне", h.Age, "лет.")
// }

func main() {
	var s Speaker //переменная интерфейсного типа
	// s = Human{"Анна", 18}
	s = Human{"Анна"}
	s.Speak()
	// s.MyAge()

	s = Dog{"Шарик"}
	s.Speak()
}
