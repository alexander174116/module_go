package main

import "fmt"

func main() {
	// counter := 1
	/*for i := 1; i <= 10; { 		//cycle like counter
		counter += 1
		fmt.Println(counter)
	}

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	for counter != 0 { 				//form of cycle, while analogy
		fmt.Println(counter)
	}

	for { 							//infinitie cycle, for stop - need 'break'
		fmt.Println("asd")
	}
	for {
		var input string  			//reference
		fmt.Print("Введите 'стоп' для выхода: ")
		fmt.Scanln(&input)

		if input == "стоп" || input == "stop" {
			break
		}
		fmt.Println("Вы ввели: ", input)
	}

	for i := 1; i <= 10; i++ { 		//reference with 'continue'
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	var age int
	fmt.Print("Введите ваш возраст: ")
	fmt.Scan(&age)
	if age < 18 {
		fmt.Println("несовершеннолетний")
	} else if age <= 64 {
		fmt.Println("совершеннолетний")
	} else {
		println("пенсионер")
	}

	var age int
	fmt.Print("Введите ваш возраст: ")
	// fmt.Scan(&age)
	if fmt.Scan(&age); age < 18 {
		fmt.Println("несовершеннолетний")
	} else if 18 <= age && age <= 64 {
		fmt.Println("совершеннолетний")
	} else {
		println("пенсионер")
	}

	sum := 0
	fmt.Print("Введите число: ")
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
		// fmt.Println(sum)
	}
	fmt.Println(sum)

	   score := 82
	   switch {
	   case score >= 90:

	   	fmt.Println("Оценка: А")

	   case score >= 75:

	   	fmt.Println("Оценка: B")

	   case score >= 60:

	   	fmt.Println("Оценка: C")

	   default:

	   		fmt.Println("Оценка: F")
	   	}
	var age int
	for i := 0; i < 5; i++ {
		fmt.Print("Введите возраст: ")
		fmt.Scan(&age)
		if age < 0 {
			fmt.Println("Ошибка: возраст не может быть отрицательным")
			continue
		} else if age < 18 {
			fmt.Println("Ребенок")
		} else if age < 66 {
			fmt.Println("Взрослый")
		} else if age >= 66 {
			fmt.Println("Пенсионер")
		}
		switch {
		case age < 18:
			fmt.Println("Ребенок")
		case age <= 65:
			fmt.Println("Взрослый")
		case age >= 66:
			fmt.Println("Пенсионер")
		}
	}*/
	var n, sum int
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
