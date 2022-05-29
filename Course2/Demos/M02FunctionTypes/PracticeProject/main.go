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

//example of a function that take variable number of arguments
// note that the "..." or "ellipsis" are treated as a slice inside the function
// also called as a variadic function means (it take a var number of arguments)
func getMax(vals ...int) int {
	max := -1

	for _, v := range vals {
		if v > max {
			max = v
		}

	}
	return max
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

	//calling a function that take a variable number of arguments
	fmt.Println("Max of function variable arguments")
	fmt.Println(getMax(1, 2, 3, 55, 53234, 22))

	// we can pass a slice as well to the variadic function ex:

	vSlice := []int{50, 60, 70, 2, 4444, 222, 5, 5353535, 22, 23, 455555555}
	fmt.Println(getMax(vSlice...))

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

	--------------------------------------------------------------------------------------------------
	Colsure is => a Function + Functions's Enviroment
	--------------------------------------------------------------------------------------------------
	==============================> Variable number of arguments <====================================
	we can define that we need a variable number of arguments using :
	(ellipsis => "..."  (three dots in a row))
	--------------------------------------------------------------------------------------------------

	Arguments of deferred calls are evaluated immediately ex:

	i := 1
	defer  fmt.Println(i + 1) // this prints 2 not 3
	i++
	fmt.Println("hello")

	output:

	hello
	2






*/
