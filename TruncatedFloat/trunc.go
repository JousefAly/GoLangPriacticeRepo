package main

import "fmt"
import "strings"
func main(){
	var inputString, lowerString string
	fmt.Println("Enter a word to test")
	fmt.Scan(&inputString)
	lowerString = strings.ToLower(inputString)
	if strings.Index(lowerString, "i") != 0 || strings.LastIndex(lowerString, "n") != len(lowerString) - 1 ||
	 !strings.Contains(lowerString, "a"){
		fmt.Println("Not Found!")
	} else {
		fmt.Println("Found!")
	}

}