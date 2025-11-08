package main //объявляем главный пакет
import "fmt" // стандартный пакет для вывода текста

func main() {
	a := 1
	b := 2
	fmt.Println("Hello, Go!") // сообщ в консоль
	if b == 0 {
		fmt.Println("ERROR DIVIDe BY ZERO")
		return //выход и программы
	}
	fmt.Print(a/b, "\n")
}
