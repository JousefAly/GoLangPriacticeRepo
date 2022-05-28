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
	fmt.Println()

	arr := [3]int{1, 2, 3}
	incrementFirstElement(&arr)
	fmt.Println(arr[0])
	slc := []int{10, 20, 30}
	incrementFirstSliceElement(slc)
	fmt.Println(slc)
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

//pass array pointer
// but this is not the right way to do in Go use slice instead !!
func incrementFirstElement(a *[3]int) {
	(*a)[0] = (*a)[0] + 1
}

// we can almost use a slice instead of an array because
// when we daclare a slice it will create its backend array

//function that increments the first slice element
func incrementFirstSliceElement(s []int) {
	s[0] += 1
}
