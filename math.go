package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
	fmt.Println(Sub(10, 10))

}

func Sub(a, b int) int {
	return a - b
}

func Soma(a, b int) int {
	return a + b
}

func Mult(a, b int) int {
	return a * b
}
