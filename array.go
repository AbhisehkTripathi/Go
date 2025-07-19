package main

import "fmt"

// array is a fixed-length
func array() {
	// Basic fixed-length array
	var a [2]any = [2]any{1, "2"}
	fmt.Println("Fixed-length array a:", a)

	// Integer array
	var b [3]int = [3]int{1, 2, 3}
	fmt.Println("Integer array b:", b)

	// Slice (dynamic array)
	var age []int = []int{2, 4, 5, 6}
	fmt.Println("Slice of ages:", age)

	// Array with explicit values
	n := [3]int{10, 20, 30}
	fmt.Println("Array n:", n)

	// Let Go count length
	mk := [...]int{1, 2, 3, 4}
	fmt.Println("Array mk (auto length):", mk, "Length:", len(mk))

	// Slice with append
	scores := []int{22, 44, 55}
	scores[1] = 98
	scores = append(scores, 66)
	fmt.Println("Scores slice:", scores)

	// String slice
	name := []string{"Abhishek", "Mohit", "Als"}
	fmt.Println("Name slice:", name)
	name = append(name, "Al")
	fmt.Println("After append:", name)

	// Slicing a slice
	name = name[1:3]
	fmt.Println("Sliced name:", name)
}

// Array iteration scenario
func iterateArray(arr [5]int) {
	fmt.Println("Iterating over array:")
	for i, v := range arr {
		fmt.Printf("Index %d: Value %d\n", i, v)
	}
}

// Slice scenario: sum of elements
func sumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	fmt.Println("Sum of slice:", sum)
	return sum
}

// DSA Example: Linear search in array
func linearSearch(arr []int, key int) int {
	for i, v := range arr {
		if v == key {
			fmt.Printf("Found %d at index %d\n", key, i)
			return i
		}
	}
	fmt.Printf("%d not found\n", key)
	return -1
}

// DSA Example: Reverse an array
func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("Reversed array:", arr)
	return arr
}

// Declare without values
var x [5]int

func Maps() {

	a := map[string]int16{
		"Abhishek": 24,
		"Mohit":    24,
		"Als":      24,
	}
	fmt.Println("aName", a["Abhishek"])
}
