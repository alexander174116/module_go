// // С использованием встроенной функции make:
// m1 := make(map[int]int)

// // Или с использованием литерала отображения:
// m2 := map[int]int{
//     12: 2,   // ключ 12, значение 2
//     1:  5,   // ключ 1, значение 5
// }

// fmt.Println(m1) // вывод: map[]
// fmt.Println(m2) // вывод: map[1:5 12:2]

// var m map[int]int
// m[12] = 3
// fmt.Println(m)

// // Вывод (сработает паника):
// // panic: assignment to entry in nil map


package main
import (
	"fmt"
)

func main() {
	
}

groupCity := map[int][]string{
	10:   []string{...}, // города с населением 10-99 тыс. человек
	100:  []string{...}, // города с населением 100-999 тыс. человек
	1000: []string{...}, // города с населением 1000 тыс. человек и более
}
cityPopulation := map[string]int{
	город: население города в тысячах человек,
}

citiesToRemove := []string{}
for city := range cityPopulation{
	flag := false
	for _, groupCityName := range groupCity[100]{
		if city == groupCityName{
			flag = true
			break
		}
	}
	if !flag{
		citiesToRemove = append(citiesToRemove, city)
		// delete(cityPopulation, city)
	}
	
}
for _, city := range citiesToRemove{
	delete(cityPopulation, city)
}


/*
func main() {
	var a int
	 m := make(map[int]int)
	for i:=0; i<10; i++{
		fmt.Scan(&a)
		// v, ok := m[a]
		if _, ok := m[a]; ok{
			fmt.Printf("%d ", m[a])
			continue
		}
		m[a] = work(a)
		fmt.Printf("%d ", m[a])
		// if ok {
		// 	fmt.Println(ok, m[a])
		// 	fmt.Printf("%d ", v)
		// }else {
		// 	m[a] = work(v)
		// }
		// fmt.Printf("%d ", work(a))
	}
	// fmt.Println(m)
	
}

func main() {
	m := map[int]int{0: 0, 1: 10}
	m2, ok := m[2]
	if !ok {
		// somehow process this case
		m2 = 20
	}
	fmt.Println(m, m[0], m[1], m2)
}

func work(n int) int {
   if n > 3 {
      time.Sleep(time.Millisecond * 500)
      return n + 1
   } else {
      time.Sleep(time.Millisecond * 500)
      return n - 1
   }
}


m := map[int]int{
	12: 2,
	1:  5,
}
fmt.Println(m[12]) // 2
delete(m, 12) // Удаляем элемент с ключом 12
fmt.Println(m) // map[1:5]


// Если мы попытаемся обратиться к ключу, которого нет в отображении, то Go вернет нулевое значение соответствующего типа
m := make(map[int]int)
fmt.Println(m[12]) // 0 (по умолчанию int = 0)
delete(m, 12)
fmt.Println(m) // map[]

// проверка на пустой мап
m := map[int]int{
	1: 10,
}
if value, inMap := m[1]; inMap {
	fmt.Println(value) // 10 (ключ 1 присутствует)
}
if value, inMap := m[2]; inMap {
	fmt.Println(value) // Условие не выполнится
}
v, ok := m[3]
    if !ok {
        fmt.Println(v, ok) // обработка
}

m := map[int]int{
	1: 10,
	2: 20,
	3: 30,
}
fmt.Println(len(m)) // 3
*/