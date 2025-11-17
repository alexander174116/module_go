// var a [3]int = [3]int{1, 2, 3}
// b := [3]int{1, 2, 3}
// c := [...]int{1, 2, 3}
// d := [3]int{1: 12}

// fmt.Println(a) // [1 2 3]
// fmt.Println(b) // [1 2 3]
// fmt.Println(c) // [1 2 3]
// fmt.Println(d) // [0 12 0]

package main

import "fmt"

func main() {    
	var workArray [10]uint8
	for i := 0; i < 10; i++ {
		var a uint8
		fmt.Scan(&a)
		workArray[i] = a
	}

	for i := 0; i < 3; i++{
		var a, b uint8
		fmt.Scan(&a, &b)
		x := workArray[a]
		y := workArray[b]
		workArray[a] = y
		workArray[b] = x
	}
	for _, elem := range workArray {
		fmt.Printf("%d ", elem)
	}
	

	
}