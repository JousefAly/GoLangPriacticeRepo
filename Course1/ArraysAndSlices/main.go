package main

import "fmt"

func main() {
	fmt.Println("Hello there")

	// C# array  =>    arr int[]    Golang array  =>   arr [5]int  Golang Slice =>  s []int

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
	lettersArray := [3]string{"yousef", "aly", "mohammed"}
	lettersArray2 := [...]string{"yousef", "aly", "mohammed"}
	lettersSlice := []string{"yousef", "aly", "mohammed"}

	fmt.Println(lettersArray[0])
	fmt.Println(lettersArray2[1])
	fmt.Println(lettersSlice[2])

	//create slice with make()

	s := make([]byte, 5)
	s[0] = 250
	fmt.Println(s[0])

	//create slice with default capacity to length when capacity is ommited

	sliceDefaultCap := make([]int, 5, 100)
	sliceDefaultCap[4] = 100
	// sliceDefaultCap[10] = 111 		//this line is invalid

	fmt.Println(sliceDefaultCap[4])
	// fmt.Println(sliceDefaultCap[10]) //invalid

	//Slicing
	//example create a slice of an array "slicing an array"

	x := [3]string{"Abdo", "Aly", "Mohammed"}
	sliceXArray := x[:2]
	fmt.Println("value of sliceXArray[1]")
	fmt.Println(sliceXArray[1])

	//Modify a slice modify the original slice

	d := []string{"a", "b", "c", "d"}
	e := d[2:]
	e[1] = "z"
	fmt.Println(d[3])

	fmt.Println("s cap and length = ")
	fmt.Println(cap(s))
	fmt.Println(len(s))
	s = s[2:4]
	fmt.Println("after first slice s cap and length = ")
	fmt.Println(cap(s))
	fmt.Println(len(s))

	//expand slice length
	//grow s length
	s = s[:cap(s)]
	fmt.Println("s cap and length after grow= ")
	fmt.Println(len(s))
	fmt.Println(cap(s))

	//double the capacity of a slice

	t := make([]byte, len(s), (cap(s)*2)+1)
	fmt.Println("t len & cap")
	fmt.Println(len(t))
	fmt.Println(cap(t))

	for i := range s {
		t[i] = s[i]
	}

	s = t
	fmt.Println("s capacity after double")
	fmt.Println(cap(t))

}
