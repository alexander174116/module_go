type Animal interface {
	Speak()
	Walk()
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Мяу! Я", c.Name)
}

func (c Cat) Walk() {
	fmt.Println(c.Name, "Идет гулять")
}

//-----------------------

// иногда интерфейс может быть пустым (без методов)
// такой интерфейс принимает любое значение
// type Empty interface{}

// var _ Speaker = Human{} // Проверка соответствия типа интерфейсу с помощью компилятора