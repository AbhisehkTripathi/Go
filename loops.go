package main

import "fmt"

func Loops() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// fmt.Println("Loop value for i", i)
	// Basic Array Declaration & Initialization
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix[1][2]) // 6

	// Find Max/Min
}

// Find Max/Min

func findMax(a [5]int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func reverse(a [5]int) [5]int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func search(a [5]int, key int) int {
	for i, v := range a {
		if v == key {
			return i
		}
	}
	return -1
}
