package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	multiplyPrint(4, 5)
	fmt.Println(multiply(5,6))
}

//create a function

func multiplyPrint(x int, y int) {
	fmt.Println(x * y)
}

func multiply(x int, y int) int{
	return x * y
}

