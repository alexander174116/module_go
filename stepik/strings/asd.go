package main
import (
	"fmt"
	"strconv"
	// "math"
	// "errors"
	"os"
	"bufio"
	// "unicode"
	"strings"
	// "unicode/utf8"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.TrimSpace(s)
	asd := strings.Split(s, ";")
	a, err := strconv.ParseFloat(asd[0], 32)
	if err != nil{
		fmt.Println("OSHIBKA")
	}
	b, err := strconv.ParseFloat(asd[1], 64)
	if err != nil{
		fmt.Println("OSHIBKA")
	}
	res := a/b
	fmt.Printf("%.4f", res)
}




/*
func main() {
	stroka1 := "%^80"
	stroka2 := "hhhhh20&&&&nd"
	res := adding(stroka1, stroka2)
	fmt.Println(res)
}

func adding(s1, s2 string)(res int64){
	scan1 := []rune(s1)
	scan2 := []rune(s2)
	for i := 0; i< len(scan1); i++{
		if !unicode.IsDigit(scan1[i]){
			scan1 = append(scan1[:i], scan1[i+1:]...)
			i--
			continue
		}
	}
	for j := 0; j < len(scan2); j++{
		if !unicode.IsDigit(scan2[j]){
			scan2 = append(scan2[:j], scan2[j+1:]...)
			j--
			continue
		}
	}
	a, _ := strconv.ParseInt(string(scan1), 10, 64)
	b, _ := strconv.ParseInt(string(scan2), 10, 64)
	return a+b
}
*/