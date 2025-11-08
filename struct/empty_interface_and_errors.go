package main

import "fmt"

// interface{} //пустой интерфейс - это интерфейс, не содержащий ни одного метода
//переменная interface{} может принимать значения любого типа, позволяя писать универсальный код

// пример пустого интерфейса:
/*
func PrintAnything(v interface{}) {
	fmt.Println("Значение:", v)
}

// теперь можно вызвать rintAnything() с чем угодно:

func main() {
	PrintAnything(42)
	PrintAnything("some text")
	PrintAnything(true)
	PrintAnything([]int{1, 2, 3})
}

// go хранит значения interface{} в специальой структуре, которая содержит:
// указатель на реальное значение
// информацию о его типе

// Извлечение конкретного типа
// При работе с interface{} вы теряете информацию о конкретном типе, ее можно вернуть с помощью type assertion:
func main() {
	var x interface{} = "Privet"

	s, ok := x.(string)
	if ok {
		fmt.Println("Eto stroka", s)
	} else {
		fmt.Println("Ne stroka")
	}
}
*/

// интерфейс error. в go ошибки - это значения, которые реализуют интерфейс error:
// - это один из самых распространенных интерфейсов в go
type error interface {
	Error() string
}

// пример интерфейса error:
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Delenie na nol' nedopustimo")
	}
	return a / b, nil
}

/*
func main() {
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Resultat:", result)
}
*/

// можно создать свою структуру и реализовать метод Error():
type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Код ошибки %d: %s", e.Code, e.Msg)
}

// создаем свою ошибку:
func DoSomething() error {
	return MyError{Code: 123, Msg: "Chto-to poshlo ne tak"}
}

func main() {
	err := DoSomething()
	if err != nil {
		fmt.Println("Oshibka", err)
	}
}
