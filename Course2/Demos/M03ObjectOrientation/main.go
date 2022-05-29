package main

import "fmt"

// defint the type
type MyInt int

//create a method that return the double of that type

func (mi MyInt) Double() int {
	return int(mi * 2)
}

func main() {
	fmt.Println("Hello World paracticing classes")
	var x MyInt = 50
	fmt.Println(x.Double())
}

/*
	What is a class ?
	a class is a bunch of data associated with a bunch of methods operates on that data.
	and togther the "methods" and "data" make what we consider to be a class
	so we need a way to associate a method with some data
	and the way it is done inside Go is using  a reciver type.

	what is a receiver type ?
	it is a type  that the method is associated with.

	we can't associate a method with an existing type.
	for ex: i can't add a method to int type
	as a general rule to associate a type with a method , both of them need to be defined
	in the same package


*/
