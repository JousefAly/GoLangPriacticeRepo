package main

import "fmt"

func main() {
	fmt.Println("Hello there")

	//declare array
	var arr [5]int

	arr[2] = 3
	fmt.Println(arr[2])

	var arr2 [2]int = [2]int{1, 2}
	fmt.Println(arr2[1])

	arr3 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(arr3[2])

	//iterating through an array
	// fmt.Println("Iterate an array")

	// for i  , v range arr2 {
	// 	fmt.Printf("index %d, val %d", i, v)
	// }

	//create slice
	lettersArray := [3] string {"yousef", "aly", "mohammed"}
	lettersArray2 := [...] string {"yousef", "aly", "mohammed"}
	lettersSlice := []string {"yousef", "aly", "mohammed"}

	fmt.Println(lettersArray[0])
	fmt.Println(lettersArray2[1])
	fmt.Println(lettersSlice[2])

}
