package main
import (
	"fmt"
	"strconv"
	// "math"
	// "errors"
	// "os"
	// "bufio"
	// "unicode"
	// "strings"
	// "unicode/utf8"
)

func main() {
	a := strconv.Itoa(2020) // int -> string
	fmt.Printf("%T \n", a) // тип - string
	fmt.Println(a) // 2020
}
/*
func main() {
	var a float64 = 1.0123456789

	// 1 параметр - число для конвертации
	// fmt - форматирование
	// prec - точность (кол-во знаков после запятой)
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	fmt.Println(strconv.FormatFloat(a, 'f', 2, 64)) // 1.01

	// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
	fmt.Println(strconv.FormatFloat(a, 'f', -1, 64)) // 1.0123456789

	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	var b float64 = 2222 * 1023 * 245 * 2 * 52
	fmt.Println(strconv.FormatFloat(b, 'e', -1, 64)) // 5.791874088e+10
}

const (
	k float64 = 1296
	p float64 = 6 
	v float64 = 6
)

func main() {
	m, w, t := M(p, v), W(k), T()
	fmt.Printf("m = %f\nv = %f\nt = %f\n", m, w, t)
}

func M(p, v float64)(m float64){
	return p*v
}

func W(k float64)(w float64){
	m:= math.Sqrt(k/M(p, v))
	return m
}

func T()(t float64){
	res := 6/W(k)
	return res
}


func main() {
	var a string
	fmt.Scan(&a)
	search := []rune(a)
	max := '0'
	for i:= range(search){
		if search[i] > max{
			max = search[i]
		}
	}
	fmt.Println(string(max))
}


func main() {
	var a string
	fmt.Scan(&a)
	my_str := []rune(a)
	// my_arr:= []rune()
	my_arr := []string{}
	// my_arr = append(my_arr, string(my_str[0]))
	// my_arr = append(my_arr, "1")
	for i:=0; i<len(my_str); i++{
		if i == 0{
			// my_arr = append(my_arr, "*")
			my_arr = append(my_arr, string(my_str[i]))
		} else {
			my_arr = append(my_arr, "*")
			my_arr = append(my_arr, string(my_str[i]))
		}
		
	}
	for i:= range(my_arr){
		fmt.Printf("%v", my_arr[i])
	}
	// fmt.Println(my_arr, my_str)
}


func main() {
    fmt.Println(divide(15, 5)) // Работает нормально
    fmt.Println(divide(4, 0))  // Генерирует ошибку
    fmt.Println("Program has been finished") // Не выполнится
}

func divide(x, y float64) float64 {
    if y == 0 {
        panic("division by zero!") // Генерация ошибки
    }
    return x / y
}


func main() {
    var a, b, res int
    fmt.Scan(&a, &b)
    res, err := divide(a, b)
    if err != nil {
        fmt.Println("ошибка")
    } else {
        fmt.Println(res)
    }
}


func main(){
	var a string
	var myFlag bool
	fmt.Scan(&a)
	a_rune := []rune(a)
	for i := range a_rune{
		// if !(unicode.Is(unicode.Latin, a_rune[i]) || unicode.IsDigit(a_rune[i])){
		if !((a_rune[i] >= '0' && a_rune[i] <= '9') || (a_rune[i] >= 'A' && a_rune[i] <= 'Z') || (a_rune[i] >= 'a' && a_rune[i] <= 'z')){
			myFlag = false
			break
		} else {
			myFlag = true
		}
	}
	if myFlag && len(a) > 5{
		fmt.Println("Ok")
	}	else {
		fmt.Println("Wrong password")
	}
}


func main(){
	var a string
	fmt.Scan(&a)
	rune_a := []rune(a)
	for i := range(rune_a){
		if strings.Count(a, string(rune_a[i])) > 1{
			rune_a[i] = 'Ё'
		}
	}
	a = string(rune_a)
	b := strings.Replace(a, "Ё", "", -1)
	fmt.Println(b)
}

func main(){
	var a string
	var myFlag bool
	a, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	orig_a := []rune(a)
	orig_a = append(orig_a[:len(orig_a)-1])
	reversed_a := []rune(a)
	for i, j := 0, len(orig_a)-1; i < len(orig_a)/2; i, j = i+1, j-1 {
		reversed_a[i], reversed_a[j] = reversed_a[j], reversed_a[i]
	}
	for i := 0; i < len(orig_a); i++{
		if orig_a[i] == reversed_a[i]{
			myFlag = true
		} else {
			myFlag = false
			break
		}
	}
	if myFlag {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}



func main(){
	var a string
	// fmt.Scanf("%s", &a) считывает до первого пробела!!
	a, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	check := []rune(a)
	// check = append(check[:len(check)-1])
	// fmt.Println(string(check[len(check)-1]))
	if unicode.IsUpper(check[0]) && string(check[len(check)-1]) == "."{
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}


func ExampleRune() {
	// Преобразуем строку в срез рун
	rs := []rune("Это срез рун")
	fmt.Printf("Оригинал: %s\n", string(rs))

	// Итерируемся по срезу рун и заменяем символ 'р' на '*'
	for i := range rs {
		if rs[i] == 'р' {
			rs[i] = '*'
		}
	}
	
	// Выводим изменённый срез в виде строки
	fmt.Printf("Изменённый срез в виде строки: %s\n", string(rs))

	// Output:
	// Изменённый срез в виде строки: Это с*ез *ун
}
*/
// func main() {
// 	var en = "english"
// 	var ru = "русский"
// 	fmt.Println(len(en), len(ru))
// 	fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru))
// }

// func main() {
// 	var txt string = "Слово"
// 	txtRunes := []rune(txt)
// 	fmt.Printf("\nСрез рун: %v", txtRunes)
// 	fmt.Printf("\nИз рун получаем снова текст: %v", string(txtRunes))
// 	lenTxtRunes := len(txtRunes)
// 	fmt.Printf("\nДлина текста (среза): %v", lenTxtRunes)

// 	for i, v := range txtRunes {
// 		fmt.Printf("\nСимвол %d: %s", i, string(v))
// 	}

// 	txtRunes[lenTxtRunes-1] = rune('а')
// 	fmt.Printf("\nЗаменили последнюю букву на \"а\": %v", string(txtRunes))

// 	txtRunes = append(txtRunes[:4])
// 	fmt.Printf("\nОтрезали последнюю букву: %v", string(txtRunes))

// 	txtRunes = append(txtRunes, 'а', 'р', 'ь')
// 	fmt.Printf("\nДобавили три буквы в конец: %v", string(txtRunes))
// }