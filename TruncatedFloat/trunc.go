package main

import (
	"fmt"
	// "strconv"
	"strings"
)

func main() {
	var inputString, lowerString string
	fmt.Println("Enter a word to test")
	fmt.Scan(&inputString)
	lowerString = strings.ToLower(inputString)
	if strings.Index(lowerString, "i") != 0 || strings.LastIndex(lowerString, "n") != len(lowerString)-1 ||
		!strings.Contains(lowerString, "a") {
		fmt.Println("Not Found!")
	} else {
		fmt.Println("Found!")
	}
	// practice quiz

	// i, _ := strconv.Atoi("10")
	// y := i * 2
	// fmt.Println(y)

	var x int
	var y *int
	z := 3
	y = &z
	x = &y
	fmt.Println(x)

}
