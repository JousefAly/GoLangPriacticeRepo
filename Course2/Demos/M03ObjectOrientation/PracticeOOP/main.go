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
