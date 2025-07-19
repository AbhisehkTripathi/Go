package main

import (
	"bufio"
	"fmt"
	"os"
)

// SaveUserInputToFile prompts user for input and saves it to a file
func SaveUserInputToFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to save to file: ")
	text, _ := reader.ReadString('\n')

	file, err := os.Create("user_input.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Input saved to user_input.txt")
}
