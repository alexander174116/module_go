/*
Длина (length) — количество элементов в срезе. len()
Емкость (capacity) — общее количество элементов от начала среза до конца базового массива. cap()
var a []int
var b []int = []int{2, 4, 6}
c := []int{2, 4, 6}
d := []int{2: 10}

fmt.Println(a) // []
fmt.Println(b) // [2 4 6]
fmt.Println(c) // [2 4 6]
fmt.Println(d) // [0 0 10]


make([]T, length, capacity)
При создании среза через make в памяти выделяется базовый массив, а срез получает к нему доступ. Его указатель
ссылается на первый элемент массива, а длина и емкость определяются в зависимости от переданных параметров.

func append(slice []Type, elems ...Type) []Type
Первый аргумент – это срез, в который добавляются новые элементы. Второй и последующие аргументы – элементы того же
типа, что и срез, которые необходимо добавить. append возвращает новый срез, включающий исходные элементы и добавленные значения:
numbers := []int{10, 20, 30}
numbers = append(numbers, 40, 50)
fmt.Println(numbers) // [10 20 30 40 50]

В Go отсутствует встроенная функция для удаления элемента из среза, но мы можем воспользоваться функцией append для того, чтобы
создать новый срез, включающий в себя срез элементов до игнорируемого элемента, а также все элементы после игнорируемого:
a := []int{1, 2, 3, 4, 5, 6, 7}
a = append(a[0:2], a[3:]...)
fmt.Println(a) // [1 2 4 5 6 7]

func ExampleExpandSlice1() {
    s := []interface{}{1, 2, 3, 4, 5}

    fmt.Println(s)   // Выводит: [1 2 3 4 5]
    fmt.Println(s...) // Выводит: 1 2 3 4 5
}

func ExampleExpandSlice2() {
    s1 := []int{1, 2, 3, 4, 5}
    s2 := []int{6, 7, 8, 9, 10}

    // Функция append(slice []Type, elems ...Type) []Type

    // Простой append не сработает, потому что второй аргумент должен быть типом int.
    s1 = append(s1, s2...)
    fmt.Println(s1)

    // Output:
    // [1 2 3 4 5 6 7 8 9 10]
}

func myPrint(a ...interface{}) {
    for _, elem := range a {
        fmt.Printf("%v ", elem)
    }
}

func ExampleMyPrint() {
    myPrint(1, "Hello", 3.14, true)

    // Output:
    // 1 Hello 3.14 true
}
*/

// package main
// import "fmt"

// func main()  {
// 	var n int
// 	fmt.Scan(&n)
// 	fmt.Println((n-1)%9 + 1)
// }

package main

import "fmt"

func main() {
	a, b := sumInt(1, 0, 123, 1,1,2,3)
	fmt.Println(a, b)
}

func sumInt(n ...int)(cnt, sum int){
	for _, value := range(n){
		sum += value
	}
	return len(n), sum
}

// func vote(x int, y int, z int) int {
// 	if x == y || x == z {
// 		return x
// 	} else if y == z {
// 		return y
// 	}
// 	return z
// }

// func minimumFromFour() int {
//     var a, min int

//     for i:=0; i<4; i++{
//         fmt.Scan(&a)
//         if i == 0{
//             min = a
//         } else if min > a{
// 			min = a
//         }
//     }
// 	return min
// }
