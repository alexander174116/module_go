package main
import "fmt"

func main() {
		x1 := 2
		x2 := 4
        test(&x1, &x2)
// 		Иначе можно записать так:
// 		a := 200
// 		var b *int = &a
}



func test(x1 *int, x2 *int) {
	a:= *x1
	b:= *x2
	*x1 = b
	*x2 = a
	fmt.Println(*x1, *x2)
}
