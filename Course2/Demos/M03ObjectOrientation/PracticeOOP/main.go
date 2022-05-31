package main

import (
	"fmt"

	"strconv"
)

// create a struct
type Point struct {
	x float64
	y float64
}

type Employee struct {
	name   string
	salary float64
}

func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}

/*
	why *Employee ? to pass by reference and not by value ok ? ok .
	why not using dereferencing to modify the actual value ex : *e.name = name  or *e.salary = salary ?
	because Golang compilere make it "Automatically". so no need to do this.
	in general: Dereferencing is automatic with . operator.
	also no need for reference when call the mehtod func (e *Employee) InitMe(name string, salary float64){}
	no need to pass the address to the pointer by making &emp.InitMe(...) just make it emp.InitMe().
	why ? because referencing is happening automatically as well like dereferencing
	// thanks Golang

	Using Pointer Receivers
	it is a good practice to make all methods of a type  to have pointer reciever OR
	make all methods with non-pointer recievers because it is easy to get confused when some are using pointer receivers
	and some are not.


*/

/*
	Golang Does not have inheritance. Golang doesn't have inheritance

*/
func (e *Employee) InitMe(name string, salary float64) {
	e.name = name
	e.salary = salary
}

func main() {
	fmt.Println("Hello world!")

	var point Point

	point.InitMe(50, 60)

	fmt.Println(point.x)
	fmt.Println(point.y)

	var employee Employee
	employee.InitMe("Yousef", 10000.32)

	fmt.Println("Hello " + employee.name)
	fmt.Println("Your slary is: ", employee.salary)
	fmt.Println("Your slary is: (other float format)" + strconv.FormatFloat(employee.salary, 'E', -1, 64))

	e := Employee{"Abdelrahman", 5000}
	fmt.Println("Hello " + e.name)
	fmt.Println("Your slary is: ", e.salary)

	e.InitMe("abdo 2", 6000)
	fmt.Println("Hello " + e.name)
	fmt.Println("Your slary is: ", e.salary)

}
