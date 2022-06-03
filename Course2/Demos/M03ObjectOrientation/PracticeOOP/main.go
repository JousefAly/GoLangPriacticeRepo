package main

import "fmt"

// create a struct
type Point struct {
	x float64
	y float64
}

func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}

func main() {
	fmt.Println("Hello world!")

	var point Point

	point.InitMe(50, 60)

	fmt.Println(point.x)
	fmt.Println(point.y)
}

