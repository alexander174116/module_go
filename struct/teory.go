// package main

import "fmt"

type Speaker interface {
	//объявляем интерфейс (набор методов)
	//интерфейс - это абстракция, описание поведения
	Speak()
}

type Dog struct {
} //объявление структуры, структура - это конкретный тип с данными

func (d Dog) Speak() {
	fmt.Println("Гав-гав!") //метод speak для Dog
}

// func main() {
	var s Speaker //переменная интерфейсного типа
	s = Dog{}     //присваиваем значение структуры
	s.Speak()     //вызов метода через интерфейс
// }

/*
//интерфейс может ссылаться на разные типы
//структура ничего не знает про интерфейсы, которые реализует

//пример интерфейса - набора методов, которые долже реализовывать тип
type Mover interface {
	Move()
}

//пример структуры - набора данные. структуры реализуют интерфейсы, если у них уже есть нужные методы
type Person struct {
	Name string
	Age ing
}

var s Speaker
s = Dog{}
s.Speak() //"Гав-гав!"

s = Person{name: "Иван"}
s.Speak() //"Привет, меня зовут Иван!"
*/

/*

//неявная реализация интерфейса
//если у структуры есть нужные метода - она автоматически считается реализацией интерфейса

type Printer interface {
	Print()
}
func SendToPrint(p Printer) {
	p.Print()
}
*/

// базовый синтаксисо бъявления интерфейса
// type InterfaceName interface {
// 	Method1(params) return_values
// 	Method2(params) return_values
// }
// примерs (у интерфейсов нет тела метода, мы не описываем, как работает метод):

// type Speaker interface {
// 	Speak()
// }

type Animal interface {
	Speak()
	Walk()
}
