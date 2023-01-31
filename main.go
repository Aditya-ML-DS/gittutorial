package main

import "fmt"

func main() {

	a := 10
	b := 20
	fmt.Println("sum of two number ", a+b)

	//  there is another change
	l := Mul(a, b)
	fmt.Println(l)

}

func Mul(a int, b int) int {
	return a * b

}
