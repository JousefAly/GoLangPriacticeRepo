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

	fmt.Println("Pass by reference example")

	z := 50
	fmt.Print("z before passing to function = ")
	fmt.Print(z)
	passByRef(&z)
	fmt.Println()
	fmt.Print("z after passing to function = ")
	fmt.Print(z)

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

/*
	Call by value
	advantage : Data encapuslation.
	disadvantage : Copying time for ex: giant slice.

*/

/*
	Call by reference.
	we can pass a pointer to the function as an argument and this is pass by reference / pointer

	advantages: copying time . we need to just copy the reference
	disadvantage: violates data encapsulation so if there is a bug inside the function it will alter the variable

*/

//pass by reference function now we call this function as "	call by reference"

func passByRef(y *int) {
	*y = *y + 1
}
