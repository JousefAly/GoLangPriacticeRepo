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
	// thanks Golang

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
}
