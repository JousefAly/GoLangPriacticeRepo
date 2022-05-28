package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	multiplyPrint(4, 5)
	fmt.Println(multiply2(5, 6))
	fmt.Println("Multiple return function will be called now")
	x, y := numWithDouble(50)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("Hahahahahahh, Ok!! nicely done Golang")
}

//create a function

func multiplyPrint(x int, y int) {
	fmt.Println(x * y)
}

func multiply(x int, y int) int {
	return x * y
}

//another way to write paramteres of the same type

func multiply2(x, y int) int {
	return x * y
}

// return multiple values

func numWithDouble(n int) (int, int) {
	return n, n * 2
}
