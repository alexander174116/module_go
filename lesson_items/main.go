package main
import "fmt"

const prog = "GoFit"

func main(){
	name := "Иван"
	age := 30
	height := 1.75
	isSub := true
	fmt.Printf("Добро пожаловать в %s!\n", prog)
	fmt.Printf("Профиль пользователя:\nИмя: %s\n", name)
	fmt.Printf("Возраст: %d лет\n", age)
	fmt.Printf("Рост: %.2f м\n", height)
	fmt.Printf("Подписан на рассылку: %t\n", isSub)
}