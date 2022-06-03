package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	//Golang doesn't have inheritance
	//we use interfaces to acheive plymorphism and do the same thing . Instructor thinks it is better way
	/*
		Interface :
		1.is a set of method signatures
		2.No implemintaion needed
		3.why we use interfaces ?
			Interfaces used to express conceptual similarity between types.

		***** in Golang we don't need to explcitly say this type implement this interface
			  Compiler will figure it automatically
			  just make the type and define it is methods. the compiler will figure theat
			  this type implement this interface.*****

		Interfaces have values (2 components):
			Dynamic type : the concrete type the variable is associated to.
			Dynamic value : the value of that concrete type.
		for example Interface Speaker{Speak()}
		and type dog that implement Speak() and have name field
		I can make var of Speaker interface like this:
		var s Speaker
		var d Dog{"Brian"}

		s = d
		s.Speak()

		so dynamic type is Dog
		and dynamic value is d




	*/
}
