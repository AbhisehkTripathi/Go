package main

import "fmt"

// array is a fixed-length
func array() {
	var a [2]any = [2]any{1, "2"}
	var b [3]int = [3]int{1, 2, 3}
	var age []int = []int{2, 4, 5, 6}
	fmt.Println("Age for Array", age)
	fmt.Println("Listing for array a", a)
	fmt.Println("Listing for array b", b)
	//  Declare with values
	n := [3]int{10, 20, 30}
	fmt.Println("Listing for array n", n)

	// Let Go count length
	mk := [...]int{1, 2, 3, 4}
	fmt.Println("Listing for array mk", len(mk))

	var scores = []int{22, 44, 55}
	scores[1] = 98
	scores = append(scores, 66)
	fmt.Println("Listing for array scores", scores)

	var name []string = []string{"Abhishek", "Mohit", "Als"}
	fmt.Println("Listing for array name", name)
	name = append(name, "Al")
	fmt.Println("Listing for array name", name)

	name = name[1:3]
	fmt.Println("Listing for final name", name)
}

// Declare without values
var x [5]int
