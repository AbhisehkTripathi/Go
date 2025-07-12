package main

import "fmt"

// array is a fixed-length
func array() {
	var a [2]any = [2]any{1, "2"}
	var b [3]int = [3]int{1, 2, 3}
	fmt.Println("Listing for array a", a)
	fmt.Println("Listing for array b", b)
	//  Declare with values
	n := [3]int{10, 20, 30}
	fmt.Println("Listing for array n", n)

	// Let Go count length
	mk := [...]int{1, 2, 3, 4}
	fmt.Println("Listing for array mk", mk)

}

// Declare without values
var x [5]int
