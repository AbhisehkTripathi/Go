package main

import (
	"errors"
	"fmt"
)

var FunctionGreeting string = "Welcome to Function in Go"

// Basic Function (No Params, No Return)
func function() {
	fmt.Println(FunctionGreeting)
}

// Function with Parameters
func greet(name string) string {
	fmt.Println("Hello,", name)
	return name
}

// func (u User) greet() string {
// 	return "Hello, " + u.Name
// }

// Function with Return Value
func add(a, b int) int {
	return a + b
}

// Function with Multiple Return Values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	} else {
		return a / b, nil
	}
}

func PrintAny[T any](value T) {
	fmt.Println(value)
}

// func getInitaal(n string) (string string) {
// 	s := strings.ToUpper(n)
// 	name := strings.Split(s, "")

// 	var initials []string
// 	for _, v := range name {
// 		initials = append(initials, v)
// 	}
// 	return strings.Join(initials, "")
// }

func CycleName[n []string](value n) {
	fmt.Println("CycleName")
	for _, v := range value {
		fmt.Println(v)
	}
}

func ConvertAny[G any](value G) G {
	return value
}

//	func checkSwith(o any) {
//		switch ys := o.(type) {
//		case int:
//			fmt.Println("int:", ys)
//		case string:
//			fmt.Println("string:", ys)
//		default:
//			fmt.Println("unknown type")
//		}
//	}
func checkType(x any) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}
}
