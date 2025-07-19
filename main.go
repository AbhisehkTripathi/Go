package main

import "fmt"

func main() {
	fmt.Println(Greeting)
	fmt.Println(FunctionGreeting)
	// fmt.Println("Age:", Age)
	fmt.Println("City:", City)
	fmt.Println("Country:", Country)
	fmt.Println("Language:", Language)
	fmt.Println("Amount:", Amount)
	fmt.Println("Email:", Email)
	variable() // Function call
	function() // Function call
	fmt.Println(add(1, 2))
	PrintAny(true)
	divide(1, 2)
	PrintAny([]int{1, 2, 3})
	checkType(true)
	CycleName([]string{"Abhishek", "Mohit", "Als"})
	array()
	Loops()
	findMax([5]int{1, 2, 3, 4, 5})
	fmt.Println(greet("888"))
	fmt.Println(VariableInfo())
	fmt.Println(CityInformation())
	fmt.Println("reverse array", reverse([5]int{1, 2, 3, 4, 5}))
	Maps()
	PointerBasics()
	ReceiverPointerDemo()
	SaveUserInputToFile()
}
