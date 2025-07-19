package main

import "fmt"

// PointerBasics demonstrates pointer usage in Go
func PointerBasics() {
	// Declare a variable
	n := 42
	fmt.Println("Value of n:", n)

	// Get the pointer (address) of n
	ptr := &n
	fmt.Println("Pointer to n (address):", ptr)
	fmt.Println("Value at pointer ptr:", *ptr)

	// Modify n via pointer
	*ptr = 100
	fmt.Println("n after modification via pointer:", n)

	// Pointer to a struct
	type Point struct {
		X int
		Y int
	}
	p := Point{X: 1, Y: 2}
	pPtr := &p
	fmt.Println("Original Point:", p)
	pPtr.X = 10
	pPtr.Y = 20
	fmt.Println("Modified Point via pointer:", *pPtr)
}
