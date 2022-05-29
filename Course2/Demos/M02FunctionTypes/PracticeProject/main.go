package main

import "fmt"

func foo(x int) int {
	return x + 1
}
func foo2(x int) int {
	return x * x
}

// apply anonymous function to an int value
func applyFuncToIntVal(xFunc func(int) int,
	val int) int {
	return xFunc(val)
}

func main() {
	fmt.Println("Hello World!")

	// declare a variable of function type
	// instead of the following line declaration we can merge the declaration with the assignment as
	// in the next line
	// var funcVar func(int) int

	funcVar := foo
	fmt.Println(funcVar(50))
	// That was nicely done by Golang actually !!!

	//apply anonymous function to an int value
	//given only the function signature so we can make this :
	var x = 50
	fmt.Println(applyFuncToIntVal(foo, x))
	fmt.Println(applyFuncToIntVal(foo2, x))

	//again Nicely done Golang!!!!!!

	//Anonymous functions are functions without a name it is somehow like lambda expression
	value := applyFuncToIntVal(func(x int) int { return x + 2 }, x)
	fmt.Println(value)

}

/*
	Golang allow us to treat functions as types
	for ex: we can define a variable of a function type,
	we can return it as a return type, we can created a function inside functions,
	can be passed as arguments

	So :
	============================> Functions are first class <===================================

	I think it is some how looke like delegates in C# when we assign a function to a var of Type function.
	Again Nicely done Golang!!!




*/
