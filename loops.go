package main

import "fmt"

// Loops demonstrates various types of loops in Go
func Loops() {
	// 1. Basic for loop (C-style)
	for i := 0; i < 5; i++ {
		fmt.Println("Basic for loop, i:", i)
	}

	// 2. for as while loop
	i := 0
	for i < 3 {
		fmt.Println("While-style for, i:", i)
		i++
	}
	for i := 0; i < 6; i++ {
		fmt.Println("While-style for, i:", i)
	}

	// i := 0
	// for i < 5 {
	// 	fmt.Println("While-style for, i:", i)
	// 	i++
	// }
	// 3. Infinite loop (use break to exit)
	// count := 0
	// for {
	// 	fmt.Println("Infinite loop, count:", count)
	// 	count++
	// 	if count >= 2 {
	// 		break
	// 	}
	// }

	// 4. Range-based loop (slice)
	slice := []string{"a", "b", "c"}
	for idx, val := range slice {
		fmt.Printf("range over slice: idx=%d, val=%s\n", idx, val)
	}

	// 5. Range-based loop (map)
	m := map[string]int{"x": 1, "y": 2}
	for k, v := range m {
		fmt.Printf("range over map: key=%s, value=%d\n", k, v)
	}

	// 6. Range-based loop (string)
	for idx, r := range "Go!" {
		fmt.Printf("range over string: idx=%d, rune=%c\n", idx, r)
	}

	// 7. Nested loops (2D array)
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("matrix[%d][%d]=%d\n", i, j, matrix[i][j])
		}
	}

	// 8. Loop control: break and continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip when i == 2
		}
		if i == 4 {
			break // Exit loop when i == 4
		}
		fmt.Println("Loop control, i:", i)
	}

	// Practical: Find max in slice
	nums := []int{3, 7, 2, 9, 1}
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Max value:", max)

	// Practical: Reverse a slice (using two indices)
	arr := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("Reversed:", arr)

	// Practical: Search for a value
	key := 3
	found := -1
	for i, v := range arr {
		if v == key {
			found = i
			break
		}
	}
	if found != -1 {
		fmt.Printf("Found %d at index %d\n", key, found)
	} else {
		fmt.Printf("%d not found\n", key)
	}
}

// --- End of loop scenarios ---
// Each section above demonstrates a different scenario for loops in Go, with comments and practical usage.

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
